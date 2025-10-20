package controller

import (
	"blog/database"
	"blog/model"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Get all the Blogs

func GetAllBlogs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(database.Blogs)
}

// Get all the Blogs -> Blogs of each Author

func GetAllBlogsOfAuthor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	var dataset = database.AuthorToBlogs
	var result []string
	for author, value := range dataset {
		if author == params["id"] {
			result = append(result, value...)
		}
	}
	json.NewEncoder(w).Encode(&result)
}

// Push the blog in the database

func PublishBlog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please enter the correct input!!!")
	}

	var blog model.Blog
	_ = json.NewDecoder(r.Body).Decode(&blog)
	blog.Id = strconv.Itoa(rand.Intn(100000))
	database.Blogs = append(database.Blogs, blog)
	author := blog.Author
	value := blog.Id
	genre := blog.Genre
	database.AuthorToBlogs[author] = append(database.AuthorToBlogs[author], value)
	database.GenreToBlogs[genre] = append(database.GenreToBlogs[genre], value)
	json.NewEncoder(w).Encode("The blog has been published!!!")
}

// Update the Blog by it's title

func UpdateBlog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, value := range database.Blogs {
		if value.Id == params["id"] {
			database.Blogs = append(database.Blogs[:index], database.Blogs[index+1:]...)
			var blog model.Blog
			_ = json.NewDecoder(r.Body).Decode(&blog)
			blog.Id = params["id"]
			database.Blogs = append(database.Blogs, blog)
			json.NewEncoder(w).Encode("The blog has been update in the database!!!")
			return
		}
	}
	json.NewEncoder(w).Encode("No blog found with the given title!!!")
}

// Delete a blog from the database

func DeleteBlog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, value := range database.Blogs {
		if value.Id == params["id"] {
			database.Blogs = append(database.Blogs[:index], database.Blogs[index+1:]...)
			// Need to delete the Id from the AuthorToBlog table too
			var name = value.Author
			for indexx, valuee := range database.AuthorToBlogs[name]{
				if valuee == params["id"]{
					database.AuthorToBlogs[name] = append(database.AuthorToBlogs[name][:indexx], database.AuthorToBlogs[name][indexx+1:]...)
					break
				}
			}
			for indexx, valuee := range database.GenreToBlogs[name]{
				if valuee == params["id"]{
					database.GenreToBlogs[name] = append(database.GenreToBlogs[name][:indexx], database.GenreToBlogs[name][indexx+1:]...)
					break
				}
			}
			json.NewEncoder(w).Encode("This blog has been deleted from the database!!!")
			return
		}
	}
	json.NewEncoder(w).Encode("No blog found with the input title, please enter the correct title!!!")
}
