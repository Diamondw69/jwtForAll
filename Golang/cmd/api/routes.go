package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func (app *application) routes() http.Handler {

	router := mux.NewRouter()
	//router.HandleFunc("/login", app.loginPageHandler).Methods("GET", "OPTIONS")

	return router
}
