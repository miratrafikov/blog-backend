package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/miratrafikov/blog_backend/internal/usecase"
)

type StartServerInput struct {
	ServerPort   string
	DatabaseName string
}

func StartApplication(input StartServerInput) {
	serverAddress := fmt.Sprintf("localhost:%s", input.ServerPort)
	http.HandleFunc("/posts", usecase.GetAllPosts)
	if err := http.ListenAndServe(serverAddress, nil); err != nil {
		log.Fatal(err)
	}
}
