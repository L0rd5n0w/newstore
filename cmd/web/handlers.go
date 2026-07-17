package main

import (
	"html/template"
	"log"
	"net/http"
	// "github.com/L0rd5n0w/newstore/internals/models"
)

func(app *application) home(w http.ResponseWriter, r *http.Request) {

	books, err := app.books.All()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	t, err := template.ParseFiles("./templates/html/home.gohtml")
	if err != nil {
		log.Print(err)
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = t.Execute(w, books)
	if err != nil {
		log.Print(err)
	}
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
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), 400)
	}

	err = app.books.Insert(
		r.PostForm.Get("title"),
		r.PostForm.Get("author"),
		r.PostForm.Get("decription"),
	)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	http.Redirect(w, r, "/", http.StatusFound)
}