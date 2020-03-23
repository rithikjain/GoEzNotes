package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/rithikjain/TodoApi/api/models"
	"github.com/rithikjain/TodoApi/api/views"
)

func RegisterUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			data := views.RegisterRequest{}
			json.NewDecoder(r.Body).Decode(&data)
			err := models.CreateUser(data.Name, data.Email, data.Password)
			if err != nil {
				w.Write([]byte("Error in creating user!"))
				return
			}
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(data)
		}
	}
}
