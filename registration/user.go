package registration

import (
	"blog/database"
	"blog/model"
	"encoding/json"
	"net/http"
)

// Get all users from the database

func GetAllUsers(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	var user []model.User

	result := database.DB.Find(&user)

	if result.Error != nil{
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(&user)
}

// To Register a user based on their name and email (unique)

func RegisterUser(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	var user model.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil{
		http.Error(w, "Invalid User Input", http.StatusBadGateway)
		return 
	}

	result := database.DB.Create(&user)

	if result.Error != nil{
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "User has been registered successfully!!!",
		"user_id": user.ID,
	})
}