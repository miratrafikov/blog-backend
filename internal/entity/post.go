package entity

import (
	"database/sql"
	"time"
)

type Post struct {
	Id        int32     `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
}

func DeserializePosts(rows *sql.Rows) []Post {
	posts := []Post{}
	for rows.Next() {
		post := Post{}
		rows.Scan(&post.Id, &post.Title, &post.Content, &post.CreatedAt)
		posts = append(posts, post)
	}

	return posts
}
