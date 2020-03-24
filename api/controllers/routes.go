package controllers

import "net/http"

func Register() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/ping", ping())
	mux.HandleFunc("/user/register", RegisterUser())
	mux.HandleFunc("/user/login", LoginUser())

	return mux
}
