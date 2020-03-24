package models

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"github.com/rithikjain/TodoApi/api/utils"
	"golang.org/x/crypto/bcrypt"
	"log"
	"os"
	"strings"
)

// Validation
func (user *User) Validate() (map[string]interface{}, bool) {
	if !strings.Contains(user.Email, "@") {
		return utils.Message(false, "MalformedEmailError"), false
	}

	if len(user.Username) <= 4 || len(user.Username) > 60 {
		return utils.Message(false, "UsernameOutOfBoundsError"), false
	}

	if len(user.Password) < 6 || len(user.Password) > 60 {
		return utils.Message(false, "PassOutOfBoundsError"), false
	}

	// Check if email already exists in the database
	temp := &User{}
	err := db.Where("email = ?", user.Email).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return utils.Message(false, "DbConnectionError"), false
	}
	if temp.Email != "" {
		return utils.Message(false, "EmailAlreadyExistsError"), false
	}

	return utils.Message(true, "ValidationSuccess"), true
}

// Register a new user
func (user *User) Create() map[string]interface{} {
	if res, valid := user.Validate(); !valid {
		return res
	}

	hashedPass, err := HashPassword(user.Password)
	if err != nil {
		log.Fatalln(err)
	}
	user.Password = hashedPass

	db.Create(user)
	if user.ID == 0 {
		return utils.Message(false, "DbConnectionError")
	}

	// JWT Token Creation
	tk := JWTToken{UserID: user.ID}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tk)
	tokenString, err := token.SignedString([]byte(os.Getenv("jwt_secret")))
	if err != nil {
		log.Fatalln(err)
	}
	user.Token = tokenString
	user.Password = ""

	response := utils.Message(true, "UserCreationSuccessful")
	response["user"] = user
	return response
}

// Login
func Login(email, password string) map[string]interface{} {
	user := &User{}
	err := db.Where("email = ?", email).First(user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return utils.Message(false, "Email not found")
		}
		return utils.Message(false, "Connection error, try again!")
	}

	if !CheckPasswordHash(password, user.Password) {
		return utils.Message(false, "Invalid Login Credentials")
	}

	// Login Successful
	// JWT Token Creation
	tk := JWTToken{UserID: user.ID}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tk)
	tokenString, err := token.SignedString([]byte(os.Getenv("jwt_secret")))
	if err != nil {
		log.Fatalln(err)
	}
	user.Token = tokenString
	user.Password = ""
	res := utils.Message(true, "Login Successful")
	res["user"] = user
	return res
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
