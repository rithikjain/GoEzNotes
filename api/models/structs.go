package models

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
)

type JWTToken struct {
	UserID uint
	jwt.StandardClaims
}

type User struct {
	gorm.Model
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `gorm:"-" json:"token"`
}

type Note struct {
	NoteID uint   `gorm:"primary_key" json:"note_id"`
	UserID uint   `json:"user_id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}
