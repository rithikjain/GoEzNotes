package controllers

import (
	"github.com/gorilla/mux"
	"github.com/rithikjain/TodoApi/api/middleware"
)

func Register() *mux.Router {
	router := mux.NewRouter()

	// JWT Middleware
	router.Use(middleware.JwtAuthentication)

	// Test Route
	router.HandleFunc("/ping", ping())

	// Auth Routes
	router.HandleFunc("/api/user/register", RegisterUser()).Methods("POST")
	router.HandleFunc("/api/user/login", LoginUser()).Methods("POST")

	// Notes Routes
	router.HandleFunc("/api/notes/new", CreateNote()).Methods("POST")
	router.HandleFunc("/api/notes", ShowAllNotes()).Methods("GET")

	return router
}
