package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/L0rd5n0w/newstore/internals/models"
)

func(app *application) home(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("./templates/html/home.gohtml")
	if err != nil {
		log.Print(err)
		http.Error(w, "Internal Server Error", 500)
	}

	err = t.Execute(w, nil)
	if err != nil {
		log.Print(err)
		http.Error(w, "Internal Server Error", 500)
	}

	w.Write([]byte("This is a restart"))
}

func(app *application) form(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./templates/html/form.gohtml")
	if err != nil {
		log.Print(err)
		http.Error(w, "Internal Server Error", 500)
	}

	err = t.Execute(w, nil)
	if err != nil {
		log.Print(err)
		http.Error(w, "Internal Server Error", 400)
	}
}

func(app *application) formHandler(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	author := r.FormValue("author")
	description := r.FormValue("description")

	bb := &models.Books{
		Title: title,
		Author: author,
		Description: description,
	}
	w.Write([]byte("Saving to database"))

	log.Print(bb)
}