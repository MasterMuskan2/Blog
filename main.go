package main

import (
	"blog/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Welcome to my Blogs!!!")

	r := router.Router()

	log.Fatal(http.ListenAndServe(":2000", r))
	
}
