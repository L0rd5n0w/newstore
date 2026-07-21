package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
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
	idStr := r.PathValue("id")

	if idStr == "" {
		http.Error(w, "Missing book ID", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Book ID", http.StatusBadRequest)
		return
	}

	bookView, err := app.books.Get(id)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	t, err := template.ParseFiles("./templates/html/view.html")
	if err != nil {
		log.Fatal(err)
		http.Error(w, err.Error(), 500)
		return
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

	title := r.PostForm.Get("title")
	if title == "" {
		http.Error(w, "Title is Empty", 500)
		return
	}

	err = app.books.Insert(
		title,
		r.PostForm.Get("author"),
		r.PostForm.Get("description"),
	)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	http.Redirect(w, r, "/", http.StatusFound)
}

func(app *application) formUpdate(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")

	if idStr == "" {
		http.Error(w, "Missing book ID", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Book ID", http.StatusBadRequest)
		return
	}

	bookView, err := app.books.Get(id)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	
	t, err := template.ParseFiles("./templates/html/updateform.html")
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), 500)
	}

	err = t.Execute(w, bookView)
	if err != nil {
		log.Print(err)
	}
}

func(app *application) formUpdateSaver(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
/*	if err != nil {
		http.Error(w, err.Error(), 400)
	}
	idStr := r.PathValue("id")

	if idStr == "" {
		http.Error(w, "Missing Book ID", http.StatusBadRequest)
		return
	}*/
	idStr := r.FormValue("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Book ID", http.StatusBadRequest)
		return
	}

	err = app.books.Update(
		r.PostForm.Get("description"),
		id,
	)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	http.Redirect(w, r, "/", http.StatusFound)
}

func(app *application) delete(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")

	if idStr == "" {
		http.Error(w, "Missing Book ID", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Book ID", http.StatusBadRequest)
		return
	}
	
	err = app.books.Delete(id)
	if err != nil {
		log.Print(err)
	}

	http.Redirect(w, r, "/", http.StatusFound)
}