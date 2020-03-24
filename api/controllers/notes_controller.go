package controllers

import (
	"encoding/json"
	"github.com/rithikjain/TodoApi/api/models"
	"github.com/rithikjain/TodoApi/api/utils"
	"net/http"
)

func CreateNote() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		note := &models.Note{}
		err := json.NewDecoder(r.Body).Decode(note)
		if err != nil {
			utils.Respond(w, utils.Message(false, err.Error()))
			return
		}
		res := note.CreateNote(r.Context())
		utils.Respond(w, res)
	}
}

func ShowAllNotes() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res := models.ShowAllNotes(r.Context())
		utils.Respond(w, res)
	}
}
