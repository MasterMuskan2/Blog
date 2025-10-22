package model

import "gorm.io/gorm"

type Blog struct {
	gorm.Model
	Title             string `json:"title"`
	Genre             string `json:"genre"`
	Author            string `json:"author"`
	Content           string `json:"content"`
	YearOfPublication int    `json:"year"`
	Likes []Like `json:"likes"`
	Comments []Comment `json:"comments"`
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

type Like struct{
	gorm.Model
	UserId uint `json:"userid"`
	BlogId uint `json:"blogid"`
}

type Comment struct{
	gorm.Model
	UserId uint `json:"userid"`
	BlogId uint `json:"blogid"`
	Content string `json:"content"`
}