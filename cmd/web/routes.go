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

	//Transactions Flow Page
	mux.Get("/", handlers.Repo.FlowBoard)
	mux.Get("/flowboard", handlers.Repo.FlowBoard)

	// transaction categories
	mux.Get("/dashboard", handlers.Repo.Home)
	mux.Get("/dashboard/tcats", handlers.Repo.TransactionCategory)
	mux.Get("/dashboard/tcats/new", handlers.Repo.TransactionCategoryNew)
	mux.Post("/dashboard/tcats/new", handlers.Repo.PostTransactionCategoryNew)
	mux.Get("/dashboard/tcats/update", handlers.Repo.TransactionCategoryUpdateGet)
	mux.Post("/dashboard/tcats/update", handlers.Repo.TransactionCategoryUpdatePost)
	mux.Post("/dashboard/tcats/delete", handlers.Repo.TransactionCategoryDelete)
	// transaction data
	mux.Get("/dashboard/tdata", handlers.Repo.TransactionsData)
	mux.Get("/dashboard/tdata/{id}", handlers.Repo.TransactionsDataDetails)
	mux.Get("/dashboard/tdata/new", handlers.Repo.TransactionsDataNew)
	mux.Post("/dashboard/tdata/new", handlers.Repo.TransactionsDataNewPost)
	mux.Post("/dashboard/tdata/delete", handlers.Repo.TransactionDataDelete)
	mux.Post("/dashboard/tdata/update", handlers.Repo.TransactionDataUpdatePost)
	mux.Get("/dashboard/tdata/update", handlers.Repo.TransactionDataUpdateGet)

	//Transactions recurence
	mux.Get("/dashboard/trecurence", handlers.Repo.TransactionRecurence)
	mux.Get("/dashboard/trecurence/new", handlers.Repo.TransactionRecurenceNew)
	mux.Post("/dashboard/trecurence/new", handlers.Repo.PostTransactionRecurenceNew)
	mux.Get("/dashboard/trecurence/update", handlers.Repo.TransactionRecurenceUpdateGet)
	mux.Post("/dashboard/trecurence/update", handlers.Repo.TransactionRecurenceUpdatePost)
	mux.Post("/dashboard/trecurence/delete", handlers.Repo.TransactionRecurenceDelete)

	//Transaction types
	mux.Get("/dashboard/ttypes", handlers.Repo.TransactionTypes)

	//Transaction log
	mux.Get("/dashboard/tlog", handlers.Repo.TransactionsLog)
	mux.Get("/dashboard/tlog/{id}", handlers.Repo.TransactionsLogDetails)
	mux.Get("/dashboard/tlog/new", handlers.Repo.TransactionsLogNew)
	mux.Post("/dashboard/tlog/new", handlers.Repo.TransactionsLogNewPost)
	return mux
}
