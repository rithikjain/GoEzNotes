package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/rithikjain/TodoApi/api/models"
	"github.com/rithikjain/TodoApi/api/views"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			data := views.RegisterRequest{}
			json.NewDecoder(r.Body).Decode(&data)
			hashPass, hashErr := HashPassword(data.Password)
			if hashErr != nil {
				log.Fatal(hashErr)
				return
			}
			err := models.CreateUser(data.Name, data.Email, hashPass)
			if err != nil {
				w.Write([]byte("Error in creating user!"))
				return
			}
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte("User Succesfully Created."))
		}
	}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
