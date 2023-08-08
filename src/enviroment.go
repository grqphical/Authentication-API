package main

import (
	"log"

	"github.com/joho/godotenv"
)

// Loads env variables in .env so they can be retrieved with os.Getenv()
func LoadDotEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
