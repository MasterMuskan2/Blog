package statistics

import (
	"blog/database"
	"blog/model"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// Get the statistics of an Author -> Number of blogs, Total number of likes, Total number of comments

func GetAuthorStats(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	author_name := params["id"]
	var numOfBlogs int64
	database.DB.Model(&model.Blog{}).Where("author = ?", author_name).Count(&numOfBlogs)
	var totalCountOfLikes int64
	database.DB.Table("likes").
    Joins("JOIN blogs ON blogs.id = likes.blog_id").Where("blogs.author = ?", author_name).Count(&totalCountOfLikes)
	var comments []string
	database.DB.Table("comments").Joins("JOIN blogs on blogs.id = comments.blog_id").Where("blogs.author = ?", author_name).Pluck("comments.content", &comments)
	response := map[string]interface{}{
		"author":            author_name,
		"total_blogs":       numOfBlogs,
		"total_likes":       totalCountOfLikes,
		"total_comments":    len(comments),
		"comment_contents":  comments,
	}
	
	json.NewEncoder(w).Encode(response)
}