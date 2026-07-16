package main

import (
	"net/http"
)

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", app.home)
	mux.HandleFunc("GET /form", app.form)
	mux.HandleFunc("POST /form/post", app.formHandler)

	return mux
}