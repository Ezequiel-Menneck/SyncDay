package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"net/http"
	"os"
	"os/signal"
	"syncday/internal/api"
)

func main() {

	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	ctx := context.Background()

	conn := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s",
		os.Getenv("SYNCDAY_DATABASE_USER"),
		os.Getenv("SYNCDAY_DATABASE_PASSWORD"),
		os.Getenv("SYNCDAY_DATABASE_HOST"),
		os.Getenv("SYNCDAY_DATABASE_PORT"),
		os.Getenv("SYNCDAY_DATABASE_NAME"),
	)

	pool, err := pgxpool.New(ctx, conn)
	if err != nil {
		panic(err)
	}

	defer pool.Close()

	if err = pool.Ping(ctx); err != nil {
		panic(err)
	}

	handler := api.NewHandler(pool)
	go func() {
		if err = http.ListenAndServe(":8080", handler); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				panic(err)
			}
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
}
