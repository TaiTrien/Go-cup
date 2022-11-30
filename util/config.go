package util

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	ApiKey string `mapstructure:"API_KEY"`
}

func LoadConfig() (config Config, err error) {
	err = godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	val := os.Getenv("API_KEY")
	config.ApiKey = val
	return
}
