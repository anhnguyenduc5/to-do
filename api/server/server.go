package server

import (
	"log"
	"to-do/config"
	"to-do/database"
	"to-do/handler"
)

func RegisterServer() {
	// Load config
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Could not load config: %v", err)
	}
	db, err := database.InitDB(cfg)
	if err != nil {
		log.Fatalf("Could not initialize database: %v", err)
	}
	// Initialize router
	r := handler.RegisterHandlerV1(db)
	// Start server
	log.Printf("Server is running on port %s", cfg.ServerPort)
	if err := r.Run(":" + cfg.ServerPort); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
