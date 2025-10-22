package controller

import (
	"blog/database"
	"blog/model"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Get all the Blogs

func GetAllBlogs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var blogs []model.Blog

	result := database.DB.Find(&blogs)

	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(blogs)
}

// Get all the Blogs -> Blogs of each Author

func GetAllBlogsOfAuthor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	author := params["id"]

	var blogs []model.Blog

	result := database.DB.Where("author = ?", author).Find(&blogs)

	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	if result.RowsAffected == 0{
		json.NewEncoder(w).Encode("No blogs found for this author!!!")
		return
	}

	json.NewEncoder(w).Encode(blogs)
}

// Get all the blogs of a praticular GENRE

func GetAllBlogsOfAGenre(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	genre := params["id"]

	var blogs []model.Blog

	result := database.DB.Where("genre = ?", genre).Find(&blogs)

	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	if result.RowsAffected == 0{
		json.NewEncoder(w).Encode("No blogs found for this genre!!!")
		return
	}

	json.NewEncoder(w).Encode(blogs)
}

// Push the blog in the database

func PublishBlog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var blog model.Blog
	if err := json.NewDecoder(r.Body).Decode(&blog); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	

	result := database.DB.Create(&blog)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode("Blog has been published successfully!")
}

// Update the Blog by it's title

func UpdateBlog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	id := params["id"]

	var existingBlog model.Blog
	if err := database.DB.First(&existingBlog, id).Error; err != nil {
		http.Error(w, "Blog not found", http.StatusNotFound)
		return
	}

	var updatedData model.Blog
	if err := json.NewDecoder(r.Body).Decode(&updatedData); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Update fields
	existingBlog.Title = updatedData.Title
	existingBlog.Genre = updatedData.Genre
	existingBlog.Author = updatedData.Author
	existingBlog.Content = updatedData.Content
	existingBlog.YearOfPublication = updatedData.YearOfPublication

	database.DB.Save(&existingBlog)

	json.NewEncoder(w).Encode("Blog updated successfully!")
}

// Delete a blog from the database

func DeleteBlog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	id := params["id"]

	// Convert id string to uint
	blogID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		http.Error(w, "Invalid blog ID", http.StatusBadRequest)
		return
	}

	result := database.DB.Delete(&model.Blog{}, blogID)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	if result.RowsAffected == 0 {
		http.Error(w, "Blog not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode("Blog deleted successfully!")
}
