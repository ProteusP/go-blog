package main

import (
	"go-blog/handlers"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.GET("/posts", handlers.GetPosts)
	router.POST("/posts", handlers.CreatePost)

	router.Run(":8080")
}
