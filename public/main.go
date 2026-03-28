package main

import (
	"log"
	"os"

	"github.com/krittawatcode/go-soldier-mvc/config"
	"github.com/krittawatcode/go-soldier-mvc/routes"
)

func main() {
	// Connect to database
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Auto migrate tables
	config.Migrate(db)

	// Setup router
	r := routes.SetupRouter(db)

	// Get port from env or default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
