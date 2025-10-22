package controller

import (
	"blog/database"
	"blog/model"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Get all the likes on a particular blog

func GetLikeCount(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	blogId := params["id"]
	var count int64
	database.DB.Model(&model.Like{}).Where("blog_id = ?", blogId).Count(&count)
	json.NewEncoder(w).Encode(map[string]int64{"total_likes": count})
}

// Add like to a blog

func AddLikeToBlog(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	blogId, _ := strconv.Atoi(params["id"])

	var like model.Like

	if err := json.NewDecoder(r.Body).Decode(&like); err != nil{
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return 
	}

	like.BlogId = uint(blogId)
	database.DB.Create(&like)
	json.NewEncoder(w).Encode("The blog has been liked by the user!!!")
}

// Get all comments of a Blog

func GetAllCommentsOfBlog(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	blogId, _ := strconv.Atoi(params["id"])
	var comments []model.Comment

	database.DB.Where("blog_id = ?", blogId).Find(&comments)
	json.NewEncoder(w).Encode(comments)
}

// Add comments to a Blog

func AddCommentToBlog(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	blogId, _ := strconv.Atoi(params["id"])

	var comment model.Comment

	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil{
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return 
	}

	comment.BlogId = uint(blogId)
	database.DB.Create(&comment)
	json.NewEncoder(w).Encode("Comment added to the blog!!!")
}

// Get all the likes of a particular User

func GetAllLikesOfUser(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	userId := params["id"]
	var count int64
	database.DB.Model(model.Like{}).Where("user_id =?", userId).Count(&count)
	json.NewEncoder(w).Encode(map[string]int64{"total_likes": count})
}

// Get all the comments of a particular user

func GetAllCommentsOfUser(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	userId := params["id"]
	var comments []string
	database.DB.Model(model.Comment{}).Where("user_id =?", userId).Pluck("content", &comments)
	json.NewEncoder(w).Encode(comments)
}