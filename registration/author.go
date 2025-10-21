package registration

import (
	"blog/database"
	"blog/model"
	"encoding/json"
	"net/http"
)

// Get all the Authors from the database

func GetAllAuthors(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	var author []model.Author

	result := database.DB.Find(&author)

	if result.Error != nil{
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(&author)
}

// Add an author to the database

func RegisterAuthor(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	var author model.Author

	if err := json.NewDecoder(r.Body).Decode(&author); err != nil{
		http.Error(w, "Invalid Author Input", http.StatusBadGateway)
		return
	}

	result := database.DB.Create(&author)

	if result.Error != nil{
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Author has been registered successfully!!!",
			"author_id": author.ID,
	})
}