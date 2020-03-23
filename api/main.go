package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rithikjain/TodoApi/api/controllers"
)

func main() {
	mux := controllers.Register()
	fmt.Println("Serving...")
	log.Fatal(http.ListenAndServe("localhost:3000", mux))
}
