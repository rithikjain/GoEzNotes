package controllers

import (
	"github.com/gorilla/mux"
)

func Register() *mux.Router {
	router := mux.NewRouter()

	// Test Route
	router.HandleFunc("/ping", ping())

	// Auth Route
	router.HandleFunc("/user/register", RegisterUser()).Methods("POST")
	router.HandleFunc("/user/login", LoginUser()).Methods("POST")

	return router
}
