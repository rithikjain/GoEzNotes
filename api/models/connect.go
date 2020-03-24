package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var db *gorm.DB

func Connect() *gorm.DB {
	con, err := gorm.Open("mysql", "root:30june@/todoapi?charset=utf8&parseTime=True")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to DB...")
	db = con
	return con
}
