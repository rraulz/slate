package main

import (
	"context"
	"log"
)

func main() {
	ctx := context.Background()
	// Initialize the application
	srv, err := InitializeApp(ctx, "host=localhost dbname=postgres user=postgres password=postgres")
	if err != nil {
		log.Fatalf("Failed to initialize app: %v", err)
	}

	// Start the server
	srv.Start()
}
