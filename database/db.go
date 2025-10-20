package database

import "blog/model"

var Blogs = []model.Blog{}

var AuthorToBlogs = make(map[string][]string)

var GenreToBlogs = make(map[string][]string)