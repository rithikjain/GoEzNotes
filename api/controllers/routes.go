package controllers

import "net/http"

func Register() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/ping", ping())
	mux.HandleFunc("/users/register", RegisterUser())

	return mux
}
