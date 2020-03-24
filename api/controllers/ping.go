package controllers

import (
	"encoding/json"
	"github.com/rithikjain/TodoApi/api/models"
	"net/http"
)

func ping() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := models.PingResponse{
			Code: http.StatusOK,
			Body: "pong",
		}
		json.NewEncoder(w).Encode(data)
	}
}
