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

	// transaction categories
	mux.Get("/dashboard/tcats", handlers.Repo.TransactionCategory)
	mux.Get("/dashboard/tcats/new", handlers.Repo.TransactionCategoryNew)
	mux.Post("/dashboard/tcats/new", handlers.Repo.PostTransactionCategoryNew)
	mux.Get("/dashboard/tcats/update", handlers.Repo.TransactionCategoryUpdateGet)
	mux.Post("/dashboard/tcats/update", handlers.Repo.TransactionCategoryUpdatePost)
	mux.Post("/dashboard/tcats/delete", handlers.Repo.TransactionCategoryDelete)

	return mux
}
