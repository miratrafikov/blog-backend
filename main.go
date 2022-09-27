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
	Id          int32
	DateCreated time.Time
	Content     string
}

type Tag struct {
	Id    int32
	Name  string
	Color string
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You just hit /blog. Great!")
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/blog", rootHandler).Methods("GET")
	http.Handle("/", router)
	fmt.Println("Server starting...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
