package main

import (
	"flag"
	"log"
	"os"

	"github.com/miratrafikov/blog_backend/internal/config"
	"github.com/miratrafikov/blog_backend/internal/server"
)

func main() {
	envFileDirectory := getEnvFileDirectoryFromArgs()
	if envFileDirectory == "" {
		log.Fatal("env.yaml file directory must be provided\nHelp: -h")
	}

	config := config.ReadConfigFromEnvFile(envFileDirectory)
	server.StartServer(config.Server.Port)
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
	flagEnvFile := flag.String("f", "", "directory containing env.yaml")
	flag.Parse()

	return *flagHelp, *flagEnvFile
}
