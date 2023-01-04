package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/miratrafikov/blog_backend/internal/controller"
)

type StartServerInput struct {
	ServerPort   string
	DatabaseName string
}

func StartApplication(input StartServerInput) {
	assignControllers()
	serverAddress := fmt.Sprintf("localhost:%s", input.ServerPort)
	if err := http.ListenAndServe(serverAddress, nil); err != nil {
		log.Fatal(err)
	}
}

func assignControllers() {
	http.HandleFunc("/posts", controller.RoutePosts)
}
