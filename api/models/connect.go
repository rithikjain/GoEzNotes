package models

import (
	"database/sql"
	"fmt"
	"log"
)

var con *sql.DB

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:30june@/todoapi")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to DB...")
	con = db
	return db
}
