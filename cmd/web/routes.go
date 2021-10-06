package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/piotrzalecki/budget/internal/handlers"
)

func routes() http.Handler {
	mux := chi.NewRouter()

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	mux.Get("/dashboard", handlers.Repo.Home)

	mux.Get("/login", handlers.Repo.Login)
	mux.Get("/", handlers.Repo.Login)
	mux.Post("/login", handlers.Repo.LoginPost)
	mux.Post("/", handlers.Repo.LoginPost)

	return mux
}
