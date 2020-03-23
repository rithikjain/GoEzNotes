package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var con *gorm.DB

func Connect() *gorm.DB {
	db, err := gorm.Open("mysql", "root:30june@/todoapi")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to DB...")
	con = db
	return db
}
