package server

import (
	"fmt"
	"log"
	"net/http"
)

func StartServer(port string) {
	serverAddress := fmt.Sprintf("localhost:%s", port)
	fmt.Printf("Starting server http://%s/\n", serverAddress)
	if err := http.ListenAndServe(serverAddress, nil); err != nil {
		log.Fatal(err)
	}
}
