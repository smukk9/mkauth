package main

import (
	"log"

	"github.com/smukk9/mkauth/internal/config"
	"github.com/smukk9/mkauth/internal/db"
	"github.com/smukk9/mkauth/internal/server"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize database
	database, err := db.New(cfg.Database.Path)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer database.Close()

	// Initialize server
	srv, err := server.New(cfg, database)
	if err != nil {
		log.Fatalf("Failed to initialize server: %v", err)
	}

	// Start server
	log.Fatal(srv.Start())
}
