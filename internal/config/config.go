package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	wd, _ := os.Getwd()
	err := godotenv.Load(filepath.Join(wd, ".env"))
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	fmt.Println("JWT_SECRET:", os.Getenv("JWT_SECRET"))
}
