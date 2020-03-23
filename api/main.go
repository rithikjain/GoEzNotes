package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rithikjain/TodoApi/api/controllers"
	"github.com/rithikjain/TodoApi/api/models"
)

func main() {
	mux := controllers.Register()
	db := models.Connect()
	defer db.Close()
	fmt.Println("Serving...")
	log.Fatal(http.ListenAndServe("localhost:3000", mux))
}
