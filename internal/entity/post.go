package entity

import "time"

type Post struct {
	Id          int32     `json:"id"`
	DateCreated time.Time `json:"dateCreated"`
	Title       string    `json:"title"`
	Content     string    `json:"content"`
}
