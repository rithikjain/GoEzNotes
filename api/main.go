package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/rithikjain/TodoApi/api/controllers"
	"github.com/rithikjain/TodoApi/api/models"
)

func main() {
	mux := controllers.Register()
	db := models.Connect()
	if !db.HasTable(&models.User{}) {
		db.CreateTable(&models.User{})
	}
	defer db.Close()
	fmt.Println("Serving...")
	log.Fatal(http.ListenAndServe("localhost:3000", mux))
}
