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
		username     string `mapstructure:"username"`
		password     string `mapstructure:"username"`
		databaseName string `mapstructure:"databaseName"`
	}
}

func ReadConfigFromDotEnv(directoryPath string) Config {
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
