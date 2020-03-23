package controllers

import (
	"encoding/json"
	"github.com/rithikjain/TodoApi/api/models"
	"github.com/rithikjain/TodoApi/api/views"
	"log"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type JWTToken struct {
	AccessToken string    `json:"token"`
	ExpiresAt   time.Time `json:"expires_at"`
}

func RegisterUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			data := views.RegisterRequest{}
			_ = json.NewDecoder(r.Body).Decode(&data)
			hashPass, hashErr := HashPassword(data.Password)
			if hashErr != nil {
				log.Fatal(hashErr)
				return
			}
			id, err := models.CreateUser(data.Name, data.Email, hashPass)
			if err != nil {
				_, _ = w.Write([]byte("Error in creating user!"))
				return
			}
			w.WriteHeader(http.StatusCreated)
			user := views.User{
				UserID:    id,
				Username:  data.Name,
				Email:     data.Email,
				Password:  hashPass,
				CreatedAt: time.Now(),
			}
			_ = json.NewEncoder(w).Encode(user)
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
