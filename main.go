package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Config struct {
	SERVER_PORT string
}

type Post struct {
	Id          int32     `json:"id"`
	DateCreated time.Time `json:"dateCreated"`
	Title       string    `json:"title"`
	Content     string    `json:"content"`
}

type Tag struct {
	Id    int32  `json:"id"`
	Name  string `json:"name"`
	Color string `json:"color"`
}

var posts = []Post{
	{
		Id:          1,
		DateCreated: time.Now(),
		Title:       "The Simplest Entry",
		Content:     "One paragraph. Nothing more.",
	},
}

func getAllPosts() []Post {
	return posts
}

func postsHandler(w http.ResponseWriter, r *http.Request) {
	posts := getAllPosts()
	json.NewEncoder(w).Encode(posts)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/posts", postsHandler).Methods("GET")
	http.Handle("/", router)
	fmt.Println("Server starting...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
