package model

type Blog struct{
	Id string `json:"id"`
	Title string `json:"title"`
	Genre string `json:"genre"`
	Author string `json:"author"`
	Content string `json:"content"`
	YearOfPublication int `json:"year"`
}