package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/rithikjain/TodoApi/api/controllers"
	"github.com/rithikjain/TodoApi/api/models"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := controllers.Register()
	db := models.Connect()
	defer db.Close()
	fmt.Println("Serving...")
	log.Fatal(http.ListenAndServe("localhost:3000", router))
}
