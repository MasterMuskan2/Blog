package model

type Blog struct{
	Title string `json:"title"`
	Genre string `json:"genre"`
	Author string `json:"author"`
	Content string `json:"content"`
	YearOfPublication int `json:"year"`
}