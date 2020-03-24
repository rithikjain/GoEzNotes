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
