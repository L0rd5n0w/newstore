package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/L0rd5n0w/newstore/internals/models"
	_ "github.com/mattn/go-sqlite3"
)

type application struct {
	Books	*models.Books
}

func main() {

	db, err := sql.Open("sqlite3", "./app.db")
	if err != nil {
		log.Fatal(err)
	}
	log.Print(db)

	app := &application{
		Books: &models.Books{},
	}

	log.Print("Starting server on :8000")
	err = http.ListenAndServe(":8000", app.routes())
	if err != nil {
		log.Fatal(err)
	}
}