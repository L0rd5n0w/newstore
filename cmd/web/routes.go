package main

import (
	"net/http"
)

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", app.home)
	mux.HandleFunc("GET /book/view/{id}", app.view)
	mux.HandleFunc("GET /form", app.form)
	mux.HandleFunc("POST /form/post", app.formHandler)
	mux.HandleFunc("GET /form/update/{id}", app.formUpdate)
	mux.HandleFunc("POST /form/updated", app.formUpdateSaver)
	mux.HandleFunc("DELETE /book/delete", app.delete)

	return mux
}