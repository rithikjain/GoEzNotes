package controllers

import (
	"github.com/rithikjain/TodoApi/api/utils"
	"net/http"
)

func ping() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := utils.Message(true, "pong")
		utils.Respond(w, data)
	}
}
