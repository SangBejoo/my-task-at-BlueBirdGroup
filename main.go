package main

import (
	// ...existing imports...
	"database/sql"
	"log"
	"os"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/SangBejoo/parking-space-monitor/handler"
	"github.com/SangBejoo/parking-space-monitor/repository"
	"github.com/SangBejoo/parking-space-monitor/server"
	"github.com/SangBejoo/parking-space-monitor/usecase"
)

func main() {
	// Initialize application
	// ...existing code...

	// Use environment variables for DB connection
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPass, dbName)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Run migrations
	migrationSQL, err := os.ReadFile("migrations/init.sql")
	if err != nil {
		log.Fatal("Error reading migration file:", err)
	}

	_, err = db.Exec(string(migrationSQL))
	if err != nil {
		log.Fatal("Error running migrations:", err)
	}

	// Initialize repositories
	trxRepo := repository.NewTrxSupplyRepository(db)
	mapRepo := repository.NewMapPlaceRepository(db)

	// Initialize use cases
	trxUsecase := usecase.NewTrxSupplyUsecase(trxRepo, mapRepo)
	mapUsecase := usecase.NewMapPlaceUsecase(mapRepo)
	monitoringUsecase := usecase.NewMonitoringUsecase(trxRepo, mapRepo)

	// Initialize handlers
	trxHandler := handler.NewTrxSupplyHandler(trxUsecase)
	mapHandler := handler.NewMapPlaceHandler(mapUsecase)
	monitoringHandler := handler.NewMonitoringHandler(monitoringUsecase)

	// Start server with handlers
	server.StartServer(trxHandler, mapHandler, monitoringHandler)
}