package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/rithikjain/TodoApi/api/views"
)

func ping() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := views.PingResponse{
			Code: http.StatusOK,
			Body: "pong",
		}
		json.NewEncoder(w).Encode(data)
	}
}
