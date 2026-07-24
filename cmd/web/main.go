package main

import (
	"database/sql"
	"log"
	"net/http"

//	"github.com/L0rd5n0w/newstore/internals/models"
	"github.com/L0rd5n0w/newstore/internals/models/sqlite"
	_ "github.com/mattn/go-sqlite3"
)

type application struct {
	books	*sqlite.BooksModel
	user	*sqlite.UserModel
}

func main() {

	db, err := sql.Open("sqlite3", "./app.db")
	if err != nil {
		log.Fatal(err)
	}

	app := &application{
		books: &sqlite.BooksModel{
			DB: db,
		},
		user: &sqlite.UserModel{
			DB: db,
		},
	}

	log.Print("Starting server on :8000")
	err = http.ListenAndServe(":8000", app.routes())
	if err != nil {
		log.Fatal(err)
	}
}