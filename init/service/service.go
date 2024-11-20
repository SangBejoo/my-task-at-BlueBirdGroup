package service

import (
	"database/sql"
	"fmt"

	"github.com/SangBejoo/parking-space-monitor/init/config"
	"github.com/SangBejoo/parking-space-monitor/init/logger"
	"github.com/SangBejoo/parking-space-monitor/internal/handler"
	"github.com/SangBejoo/parking-space-monitor/internal/repository"
	"github.com/SangBejoo/parking-space-monitor/internal/usecase"
)

type Service struct {
	DB               *sql.DB
	TrxHandler       *handler.TrxSupplyHandler
	MapHandler       *handler.MapPlaceHandler
	MonitoringHandler *handler.MonitoringHandler
}

func NewService(cfg *config.Config) (*Service, error) {
	// Initialize DB connection
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		logger.Error("Failed to connect to database: %v", err)
		return nil, err
	}

	if err := db.Ping(); err != nil {
		logger.Error("Failed to ping database: %v", err)
		return nil, err
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

	return &Service{
		DB:               db,
		TrxHandler:       trxHandler,
		MapHandler:       mapHandler,
		MonitoringHandler: monitoringHandler,
	}, nil
}

func (s *Service) Close() {
	if s.DB != nil {
		logger.Info("Closing database connection")
		s.DB.Close()
	}
}
