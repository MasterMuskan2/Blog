package router

import (
	"blog/controller"
	"blog/registration"

	"github.com/gorilla/mux"
)

func Router() *mux.Router{
	r := mux.NewRouter()

	// APIs of Controller Functions
	r.HandleFunc("/blogs", controller.GetAllBlogs).Methods("GET")
	r.HandleFunc("/blog/{id}", controller.GetAllBlogsOfAuthor).Methods("GET")
	r.HandleFunc("/blog/genre/{id}", controller.GetAllBlogsOfAGenre).Methods("GET")
	r.HandleFunc("/blog", controller.PublishBlog).Methods("POST")
	r.HandleFunc("/blog/{id}", controller.UpdateBlog).Methods("PUT")
	r.HandleFunc("/blog/{id}", controller.DeleteBlog).Methods("DELETE")
	
	// APIs for Registration Functions

	r.HandleFunc("/users", registration.GetAllUsers).Methods("GET")
	r.HandleFunc("/user/register", registration.RegisterUser).Methods("POST")

	r.HandleFunc("/authors", registration.GetAllAuthors).Methods("GET")
	r.HandleFunc("/author/register", registration.RegisterAuthor).Methods("POST")

	return r
}