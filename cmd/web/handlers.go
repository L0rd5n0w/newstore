package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	
	"github.com/julienschmidt/httprouter"
)

func(app *application) home(w http.ResponseWriter, r *http.Request) {

	books, err := app.books.All()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	t, err := template.ParseFiles("./templates/html/home.html")
	if err != nil {
		log.Print(err)
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = t.Execute(w, map[string]any{"Books": books})
	if err != nil {
		log.Print(err)
	}
}

func(app *application) view(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	bookView, err := app.books.Get(id)
	if err != nil {
		http.Error(w, err.Error(), 400)
	}

	t, err := template.ParseFiles("./templates/html/view.html")
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), 500)
	}
	err = t.Execute(w, bookView)
	if err != nil {
		log.Print(err)
	}
}

func(app *application) form(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./templates/html/form.html")
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
		r.PostForm.Get("description"),
	)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	http.Redirect(w, r, "/", http.StatusFound)
}

func(app *application) formUpdate(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./templates/html/updateform.html")
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), 500)
	}

	err = t.Execute(w, nil)
	if err != nil {
		log.Print(err)
	}
}

func(app *application) formUpdateSaver(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), 400)
	}

	err = app.books.Update(
		r.PostForm.Get("author"),
		r.PostForm.Get("description"),
	)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	http.Redirect(w, r, "/", http.StatusFound)
}

func(app *application) delete(w http.ResponseWriter, r *http.Request) {
	params:= httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		log.Print(err)
	}
	
	err = app.books.Delete(id)
	if err != nil {
		log.Print(err)
	}

	http.Redirect(w, r, "/", http.StatusFound)
}