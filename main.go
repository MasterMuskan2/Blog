package main

import (
	"blog/database"
	"blog/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Welcome to my Blogs!!!")

	database.Connect()

	r := router.Router()

	log.Fatal(http.ListenAndServe(":2000", r))

}
