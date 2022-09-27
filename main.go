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
	Content     string    `json:"content"`
}

type Tag struct {
	Id    int32  `json:"id"`
	Name  string `json:"name"`
	Color string `json:"color"`
}

var posts = []Post{
	{
		Id: 1,
		DateCreated: time.Now(),
		Content: "Let's goo!",
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
