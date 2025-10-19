package controller

import (
	"blog/database"
	"blog/model"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// Get all the Blogs

func GetAllBlogs(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(database.Blogs)
}

// Get all the Blogs -> Blogs of each Author

func GetAllBlogsOfAuthor(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	var dataset = database.AuthorToBlogs
	var result []string
	for author, value := range dataset{
		if author == params["id"]{
			result = append(result, value...)
		}
	}
	json.NewEncoder(w).Encode(&result)
}

// Push the blog in the database

func PublishBlog(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	var blog model.Blog
	_ = json.NewDecoder(r.Body).Decode(&blog)
	database.Blogs = append(database.Blogs, blog)
	author := blog.Author
	value := blog.Content
	database.AuthorToBlogs[author]=append(database.AuthorToBlogs[author], value)
	json.NewEncoder(w).Encode(blog)
}