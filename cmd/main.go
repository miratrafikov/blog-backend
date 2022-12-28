package main

import (
	"github.com/miratrafikov/blog_backend/internal/config"
	"github.com/miratrafikov/blog_backend/internal/server"
)

func main() {
	config := config.ReadConfigFromDotEnv(".")
	server.StartServer(config.Server.Port)
}
