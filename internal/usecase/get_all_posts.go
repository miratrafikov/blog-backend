package usecase

import (
	"github.com/miratrafikov/blog_backend/internal/db"
	"github.com/miratrafikov/blog_backend/internal/entity"
)

func GetAllPosts() ([]entity.Post, error) {
	rows, err := db.Connection.Query("SELECT * FROM post")
	if err != nil {
		return []entity.Post{}, err
	}

	posts := entity.DeserializePosts(rows)

	return posts, nil
}
