package controllers

import (
	"encoding/json"
	"github.com/rithikjain/TodoApi/api/models"
	"github.com/rithikjain/TodoApi/api/utils"
	"net/http"
)

func RegisterUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := &models.User{}
		err := json.NewDecoder(r.Body).Decode(user)
		if err != nil {
			utils.Respond(w, utils.Message(false, err.Error()))
			return
		}

		response := user.Create()
		utils.Respond(w, response)
	}
}

func LoginUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := &models.User{}
		err := json.NewDecoder(r.Body).Decode(user)
		if err != nil {
			utils.Respond(w, utils.Message(false, err.Error()))
			return
		}
		res := models.Login(user.Email, user.Password)
		utils.Respond(w, res)
	}
}
