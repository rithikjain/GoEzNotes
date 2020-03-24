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

	con, err := gorm.Open(dbType, os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to DB...")
	if !con.HasTable(&User{}) {
		con.CreateTable(&User{})
	}
	if !con.HasTable(&Note{}) {
		con.CreateTable(&Note{})
	}
	db = con
	return con
}
