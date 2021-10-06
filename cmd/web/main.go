package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/piotrzalecki/budget/internal/config"
	"github.com/piotrzalecki/budget/internal/handlers"
	"github.com/piotrzalecki/budget/internal/render"
)

const portNumber = ":8080"

var app config.AppConfig

func main() {

	// server configuration
	app.UseCache = false

	// mandatory inits
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewRenderer(&app)
	log.Default()
	log.Println(fmt.Printf("Starting server on port %s ", portNumber))

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(),
	}

	err := srv.ListenAndServe()
	log.Fatal(err)

}
