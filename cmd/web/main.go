package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/piotrzalecki/budget/internal/config"
	"github.com/piotrzalecki/budget/internal/driver"
	"github.com/piotrzalecki/budget/internal/handlers"
	"github.com/piotrzalecki/budget/internal/render"
)

const portNumber = ":8080"

var app config.AppConfig

func main() {

	// server configuration
	app.UseCache = false

	// database connection initialisation
	log.Println("connecting to database")
	dbname := os.Getenv("DB_NAME")
	dbuser := os.Getenv("DB_USER")
	dbpass := os.Getenv("DB_PASSWORD")
	connectionString := fmt.Sprintf("host=127.0.0.1 port=5432 dbname=%s user=%s password=%s", dbname, dbuser, dbpass)
	db, err := driver.ConnectSQL(connectionString)
	if err != nil {
		log.Fatal("Cannnot connect to database! Dying...")
	}

	// mandatory inits
	repo := handlers.NewRepo(&app, db)
	handlers.NewHandlers(repo)

	render.NewRenderer(&app)

	log.Println(fmt.Printf("Starting server on port %s ", portNumber))

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)

}
