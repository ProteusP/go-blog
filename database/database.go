package database

import (
	"database/sql"
	"go-blog/models"

	_ "github.com/mattn/go-sqlite3"
)

func Connect() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./blog.db")
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func CreatePost(db *sql.DB, post *models.Post) error {
	result, err := db.Exec("INSERT INTO posts (title, content, author) VALUES (?,?,?)",
		post.Title,
		post.Content,
		post.Author)

	if err != nil {
		return err
	}

	id, _ := result.LastInsertId()

	post.ID = int(id)
	return nil
}
