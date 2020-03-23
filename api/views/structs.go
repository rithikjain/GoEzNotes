package views

import "time"

type PingResponse struct {
	Code int         `json:"code"`
	Body interface{} `json:"body"`
}

type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	UserID    int64  `gorm:"primary_key" json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string
	CreatedAt time.Time `json:"created_at"`
}
