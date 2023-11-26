package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// LoadEnv initializes environment variables from the .env file.
func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

// InitAppEnv initializes environment variables for the application.
func InitAppEnv() {
	WebhookURL = os.Getenv("WEBHOOK_PORT")
	Host = os.Getenv("HOST")
	Port = os.Getenv("PORT")
	if Port == "" {
		Port = "8080"
	}
}
