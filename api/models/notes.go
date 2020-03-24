package models

import (
	"context"
	"github.com/rithikjain/TodoApi/api/utils"
)

// Create or update a note
func (note *Note) CreateNote(ctx context.Context) map[string]interface{} {
	user := GetUser(ctx.Value("user").(uint))
	if user == nil {
		return utils.Message(false, "UserNotFoundError")
	}
	note.UserID = user.ID

	err := db.Save(note).Error
	if err != nil {
		return utils.Message(false, err.Error())
	}
	return utils.Message(true, "NoteSavedSuccessfully")
}
