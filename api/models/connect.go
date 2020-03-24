package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"os"
)

var db *gorm.DB

func Connect() *gorm.DB {
	dbType := os.Getenv("db_type")
	dbUser := os.Getenv("db_user")
	dbPassword := os.Getenv("db_password")
	dbName := os.Getenv("db_name")

	dbUri := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True", dbUser, dbPassword, dbName)

	con, err := gorm.Open(dbType, dbUri)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to DB...")
	if !con.HasTable(&User{}) {
		con.CreateTable(&User{})
	}
	db = con
	return con
}
