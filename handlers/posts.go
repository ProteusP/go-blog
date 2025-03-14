package handlers

import (
	"go-blog/database"
	"go-blog/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPosts(c *gin.Context) {
	db, err := database.Connect()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection failed"})
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM posts")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch posts"})
		return
	}

	defer rows.Close()

	var posts []models.Post

	for rows.Next() {
		var post models.Post
		rows.Scan(&post.ID, &post.Title, &post.Content, &post.Author, &post.CreatedAt)

		posts = append(posts, post)
	}

	c.JSON(http.StatusOK, posts)
}

func CreatePost(c *gin.Context) {
	var newPost models.Post

	if err := c.ShouldBindJSON(&newPost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	if newPost.Title == "" || newPost.Content == "" || newPost.Author == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "All fields (title, content, author) are required"})
		return
	}

	db, err := database.Connect()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection failed"})
		return
	}

	defer db.Close()

	err = database.CreatePost(db, &newPost)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create post"})
		return
	}

	c.JSON(http.StatusCreated, newPost)

}
