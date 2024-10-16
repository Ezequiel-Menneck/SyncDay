package api

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"net/http"
	"strconv"
	"syncday/internal/api/spec"
	"syncday/internal/pgstore"
)

type store interface {
	GetAllUsers(ctx context.Context) ([]pgstore.User, error)
	GetUserById(ctx context.Context, id int) (pgstore.User, error)
	CreateUser(ctx context.Context, arg pgstore.CreateUserParams) (pgstore.User, error)
	DeleteUserById(ctx context.Context, id int) error
	UpdateUserBaseSalary(ctx context.Context, arg pgstore.UpdateUserBaseSalaryParams) error
}

type API struct {
	r         *chi.Mux
	s         store
	validator *validator.Validate
}

type Error struct {
	Message string
}

func (e Error) Error() string {
	return e.Message
}

func (h API) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.r.ServeHTTP(w, r)
}

func NewHandler(pool *pgxpool.Pool) http.Handler {
	a := API{
		s:         pgstore.New(pool),
		validator: validator.New(validator.WithRequiredStructEnabled()),
	}

	r := chi.NewRouter()
	r.Use(middleware.RequestID, middleware.Recoverer, middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	SetupRoutes("/api", r, a)

	a.r = r
	return a
}

func (h API) handleGetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.s.GetAllUsers(r.Context())
	if err != nil {
		h.writeResponse(w, http.StatusInternalServerError, Error{Message: "something went wrong"})
		return
	}

	h.writeResponse(w, http.StatusOK, users)
}

func (h API) handleGetUserById(w http.ResponseWriter, r *http.Request) {
	strUserId := chi.URLParam(r, "userId")
	userId, err := strconv.Atoi(strUserId)
	if err != nil {
		h.writeResponse(w, http.StatusInternalServerError, Error{Message: "error processing userId"})
		return
	}

	user, err := h.s.GetUserById(r.Context(), userId)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			h.writeResponse(w, http.StatusNotFound, Error{Message: "user not found"})
			return
		}
		h.writeResponse(w, http.StatusInternalServerError, Error{Message: "internal server error"})
		return
	}

	h.writeResponse(w, http.StatusOK, user)
}

func (h API) handleDeleteUserById(w http.ResponseWriter, r *http.Request) {
	strUserId := chi.URLParam(r, "userId")
	userId, err := strconv.Atoi(strUserId)
	if err != nil {
		h.writeResponse(w, http.StatusInternalServerError, Error{Message: "error processing userId"})
		return
	}

	err = h.s.DeleteUserById(r.Context(), userId)
	if err != nil {
		h.writeResponse(w, http.StatusInternalServerError, Error{Message: "internal server error"})
		return
	}

	h.writeResponse(w, http.StatusNoContent, nil)
}

func (h API) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	var body spec.CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		h.writeResponse(w, http.StatusInternalServerError, Error{Message: err.Error()})
		return
	}

	if err = h.validator.Struct(body); err != nil {
		h.writeResponse(w, http.StatusInternalServerError, Error{Message: "invalid input" + err.Error()})
		return
	}

	user, err := h.s.CreateUser(r.Context(), pgstore.CreateUserParams{
		Name:       body.Name,
		LastName:   body.LastName,
		Email:      body.Email,
		Cellphone:  pgtype.Text{String: body.Cellphone, Valid: true},
		BaseSalary: pgtype.Numeric{Int: &body.BaseSalary, Valid: true},
	})
	if err != nil {
		h.writeResponse(w, http.StatusInternalServerError, Error{Message: err.Error()})
		return
	}

	h.writeResponse(w, http.StatusCreated, user)
}

func (h API) handleUpdateUserBaseSalary(w http.ResponseWriter, r *http.Request) {
	var body spec.UpdateUserBaseSalary
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		h.writeResponse(w, http.StatusBadRequest, Error{Message: "invalid JSON" + err.Error()})
		return
	}

	strUserId := chi.URLParam(r, "userId")
	userId, err := strconv.Atoi(strUserId)
	if err != nil {
		h.writeResponse(w, http.StatusInternalServerError, Error{Message: "error processing userId"})
		return
	}

	err = h.s.UpdateUserBaseSalary(r.Context(), pgstore.UpdateUserBaseSalaryParams{
		BaseSalary: pgtype.Numeric{
			Int:   &body.BaseSalary,
			Valid: true,
		},
		ID: userId,
	})
	if err != nil {
		h.writeResponse(w, http.StatusInternalServerError, Error{Message: "something went wrong"})
		return
	}

	h.writeResponse(w, http.StatusOK, nil)
}

func (h API) writeResponse(w http.ResponseWriter, code int, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if body != nil {
		_ = json.NewEncoder(w).Encode(body)
	}
}
