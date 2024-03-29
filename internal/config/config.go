package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port string `mapstructure:"port"`
	}
	Db struct {
		DatabaseName string `mapstructure:"databaseName"`
	}
}

func ReadConfigFromEnvFile(directoryPath string) Config {
	viper.AddConfigPath(directoryPath)
	viper.SetConfigName("env")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	config := Config{}
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal(err)
	}

	return config
}
