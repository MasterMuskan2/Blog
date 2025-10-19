package router

import (
	"blog/controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router{
	r := mux.NewRouter()

	r.HandleFunc("/blogs", controller.GetAllBlogs).Methods("GET")
	r.HandleFunc("/blog/{id}", controller.GetAllBlogsOfAuthor).Methods("GET")
	r.HandleFunc("/blog", controller.PublishBlog).Methods("POST")

	return r
}