package models

import (
	"github.com/rithikjain/TodoApi/api/views"
	"time"
)

func CreateUser(username, email, password string) (int64, error) {
	user := views.User{
		Username:  username,
		Email:     email,
		Password:  password,
		CreatedAt: time.Now(),
	}
	err := con.Create(&user).Error
	if err != nil {
		return -1, err
	}

	return user.UserID, err
}
