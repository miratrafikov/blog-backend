package server

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/miratrafikov/blog-backend/internal/controller"
)

type StartServerInput struct {
	ServerPort   string
	DatabaseName string
}

type routeMapping struct {
	Prefix     string
	Controller func(http.ResponseWriter, *http.Request)
}

var routeMappings = []routeMapping{
	{
		Prefix:     "/posts/",
		Controller: controller.RoutePosts,
	},
}

func StartApplication(input StartServerInput) {
	http.HandleFunc("/", route)
	serverAddress := fmt.Sprintf("localhost:%s", input.ServerPort)
	if err := http.ListenAndServe(serverAddress, nil); err != nil {
		log.Fatal(err)
	}
}

func route(w http.ResponseWriter, r *http.Request) {
	url := getUrl(r)
	for _, mapping := range routeMappings {
		if strings.HasPrefix(url, mapping.Prefix) {
			mapping.Controller(w, r)

			return
		}
	}

	http.NotFound(w, r)
}

func getUrl(r *http.Request) string {
	url := r.URL.String()
	url = ensureEndsWithSlash(url)

	return url
}

func ensureEndsWithSlash(url string) string {
	if url[len(url)-1:] != "/" {
		return url + "/"
	}

	return url
}
