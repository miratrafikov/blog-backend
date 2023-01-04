package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/miratrafikov/blog-backend/internal/config"
	"github.com/miratrafikov/blog-backend/internal/db"
	"github.com/miratrafikov/blog-backend/internal/server"
)

func main() {
	envFileDirectory := getEnvFileDirectoryFromArgs()
	config := config.ReadConfigFromEnvFile(envFileDirectory)
	db.EstablishConnection(config.Db.DatabaseName)
	startServer(config)
}

func getEnvFileDirectoryFromArgs() string {
	flagHelp, flagFile := readFlags()
	if flagHelp {
		flag.PrintDefaults()
		os.Exit(0)
	}

	return flagFile
}

func readFlags() (bool, string) {
	flagHelp := flag.Bool("h", false, "show help")
	flagEnvFile := flag.String("f", ".", "directory containing env.yaml")
	flag.Parse()

	return *flagHelp, *flagEnvFile
}

func startServer(config config.Config) {
	fmt.Printf("Starting server at http://localhost:%s/...\n", config.Server.Port)
	server.StartApplication(server.StartServerInput{
		ServerPort:   config.Server.Port,
		DatabaseName: config.Db.DatabaseName,
	})
}
