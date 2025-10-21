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

type User struct{
	gorm.Model
	UserName string `json:"username"`
	UserEmail string `gorm:"unique" json:"useremail"`
}

type Author struct{
	gorm.Model
	AuthorName string `json:"authorname"`
	AuthorEmail string `json:"authoremail"`
}
