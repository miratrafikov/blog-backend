package main

import (
	"time"
	"log"
	"fmt"
	"net/http"
)

type Config struct {
	SERVER_PORT string
}

type Post struct {
	Id          int32
	DateCreated time.Time
	Content     string
}

type Tag struct {
	Id    int32
	Name  string
	Color string
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You just hit /. Great!")
}

func main() {
	fmt.Println("Server started. Congratulations!")
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
