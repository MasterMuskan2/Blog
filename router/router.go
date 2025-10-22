package router

import (
	"blog/controller"
	"blog/registration"
	"blog/statistics"

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

	// APIs for Interaction Functions

	r.HandleFunc("/likes/{id}", controller.GetLikeCount).Methods("GET")
	r.HandleFunc("/like/{id}", controller.AddLikeToBlog).Methods("POST")  

	r.HandleFunc("/comments/{id}", controller.GetAllCommentsOfBlog).Methods("GET")
	r.HandleFunc("/comment/{id}", controller.AddCommentToBlog).Methods("POST")

	r.HandleFunc("/like/count/{id}", controller.GetAllLikesOfUser).Methods("GET")
	r.HandleFunc("/comment/count/{id}", controller.GetAllCommentsOfUser).Methods("GET")

	// APIs for Statistics

	r.HandleFunc("/stats/author/{id}", statistics.GetAuthorStats).Methods("GET")

	r.HandleFunc("/stats/user/{id}", statistics.GetUserStats).Methods("GET")

	return r
}