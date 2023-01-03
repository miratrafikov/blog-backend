package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/miratrafikov/blog_backend/internal/config"
	"github.com/miratrafikov/blog_backend/internal/db"
	"github.com/miratrafikov/blog_backend/internal/server"
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

	if flagFile == "" {
		log.Fatal("env.yaml file directory must be provided\nHelp: -h")
	}

	return flagFile
}

func readFlags() (bool, string) {
	flagHelp := flag.Bool("h", false, "show help")
	flagEnvFile := flag.String("f", "", "directory containing env.yaml")
	flag.Parse()

	return *flagHelp, *flagEnvFile
}

func startServer(config config.Config) {
	fmt.Printf("Starting server http://localhost:%s/\n", config.Server.Port)
	server.StartApplication(server.StartServerInput{
		ServerPort:   config.Server.Port,
		DatabaseName: config.Db.DatabaseName,
	})
}
