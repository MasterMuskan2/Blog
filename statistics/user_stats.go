package statistics

import (
	"blog/database"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// Get the statistics of an User -> Likes and Comments

func GetUserStats(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	userId := params["id"]
	var likes []string
	database.DB.Table("likes").Joins("JOIN blogs ON blogs.id = Likes.blog_id").Where("likes.user_id = ?", userId).Pluck("blogs.title", &likes)
	
	var comments []string
	database.DB.Table("comments").Joins("JOIN blogs ON blogs.id = comments.blog_id").Where("comments.user_id = ?", userId).Pluck("comments.content", &comments)
	
	response := map[string]interface{}{
		"user_id":            userId,
		"liked_blogs":		  likes,
		"user_comments":	  comments,
	}
	
	json.NewEncoder(w).Encode(response)
}