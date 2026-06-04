package main

import (
	"go-tickets/internal/config"
	"go-tickets/internal/server"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()
	// Connect to the database
	db := config.ConnectDatabase(cfg)
	// Start the server
	server.Start(cfg, db)

}
