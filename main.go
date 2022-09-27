package main

import (
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

func blogsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You just hit /blog. Great!")
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/blogs", blogsHandler).Methods("GET")
	http.Handle("/", router)
	fmt.Println("Server starting...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
