package main

import (
	"PalaemonBlog/routes"
)

func main() {
	// DB init
	//model.InitDB()

	routes.InitRouter()
}
