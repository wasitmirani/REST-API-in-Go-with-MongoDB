package main

import (
	"context"
	"crud-with-mongodb/internal/config"
	"crud-with-mongodb/internal/routes"
	"log"
	"time"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Initialize MongoDB connection
	db, err := config.ConnectDB(cfg.MongoURI)
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
	}
		// Create a root context
	ctx := context.Background()
	// Create a context with timeout for the disconnect operation
	disconnectCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// Properly disconnect from MongoDB when main exits
	defer func() {
		if err := db.Disconnect(disconnectCtx); err != nil {
			log.Printf("Warning: error disconnecting from MongoDB: %v", err)
		} else {
			log.Println("Disconnected from MongoDB")
		}
	}()

	// Setup routes
	router := routes.SetupRouter(db, cfg)

	// Start server
	log.Printf("Server starting on port %s...", cfg.Port)
	if err := router.Run(":" + cfg.Port); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}