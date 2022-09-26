package main

import (
	"fmt"
	"time"
)

type Config struct {
	SERVER_PORT string
}

type Post struct {
	Id int32
	DateCreated time.Time
	Content string
}

type Tag struct {
	Id int32
	Name string
	Color string
}

func main() {
	fmt.Printf("You just ran main.go. Good for you!")
}