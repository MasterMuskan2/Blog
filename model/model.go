package model

import "gorm.io/gorm"

type Blog struct {
	gorm.Model
	Title             string `json:"title"`
	Genre             string `json:"genre"`
	Author            string `json:"author"`
	Content           string `json:"content"`
	YearOfPublication int    `json:"year"`
}
