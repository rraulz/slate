package main

import (
	"context"
	"log"
)

func main() {
	ctx := context.Background()
	// Initialize the application
	srv, err := InitializeApp(ctx, "")
	if err != nil {
		log.Fatalf("Failed to initialize app: %v", err)
	}

	// Start the server
	srv.Start()
}
