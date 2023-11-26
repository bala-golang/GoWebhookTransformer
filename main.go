package main

import (
	"fmt"
	"log"

	"github.com/bala-golang/GoWebhookTransformer/api/handlers"
	"github.com/bala-golang/GoWebhookTransformer/api/server"
	"github.com/bala-golang/GoWebhookTransformer/config"
	"github.com/bala-golang/GoWebhookTransformer/internal"
)

func init() {
	// Initialize application environment variables
	config.InitAppEnv()
}

func main() {
	// Start worker goroutine for handling webhook requests
	go internal.Worker(handlers.RequestChannel)

	// Create and run the server
	router := server.NewServer()
	address := fmt.Sprintf(":%s", config.Port)

	log.Printf("Server is running on %s", address)
	if err := router.Run(address); err != nil {
		log.Fatal(err)
	}
}
