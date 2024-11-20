package main

import (
	"os"

	_ "github.com/lib/pq"
	"github.com/SangBejoo/parking-space-monitor/init/config"
	"github.com/SangBejoo/parking-space-monitor/init/logger"
	"github.com/SangBejoo/parking-space-monitor/init/server"
	"github.com/SangBejoo/parking-space-monitor/init/service"
)

func main() {
	// Initialize logger
	logger.Init()

	// Load configuration
	cfg := config.LoadConfig()

	// Initialize service
	svc, err := service.NewService(cfg)
	if err != nil {
		logger.Fatal("Failed to initialize service: %v", err)
	}
	defer svc.Close()

	// Run migrations
	migrationSQL, err := os.ReadFile("migrations/init.sql")
	if err != nil {
		logger.Fatal("Error reading migration file: %v", err)
	}

	_, err = svc.DB.Exec(string(migrationSQL))
	if err != nil {
		logger.Fatal("Error running migrations: %v", err)
	}

	// Start server with handlers from service
	server.StartServer(svc.TrxHandler, svc.MapHandler, svc.MonitoringHandler)
}