package config

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	// wd, _ := os.Getwd()
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
