package main

import (
	"PalaemonBlog/model"
	"PalaemonBlog/routes"
)

func main() {
	// DB init
	model.InintDB()

	routes.InitRouter()
}
