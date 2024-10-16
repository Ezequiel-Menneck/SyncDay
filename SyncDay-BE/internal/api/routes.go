package api

import (
	"github.com/go-chi/chi/v5"
)

func SetupRoutes(baseUrl string, r chi.Router, a API) *chi.Router {
	r.Route(baseUrl, func(r chi.Router) {
		r.Route("/users", func(r chi.Router) {
			r.Get("/", a.handleGetAllUsers)
			r.Get("/{userId}", a.handleGetUserById)
			r.Post("/", a.handleCreateUser)
			r.Delete("/{userId}", a.handleDeleteUserById)
			r.Put("/{userId}", a.handleUpdateUserBaseSalary)
		})
	})

	return &r
}
