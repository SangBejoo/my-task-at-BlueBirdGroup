package usecase

import (
	"context"
	"os"
	"fmt"
	"github.com/SangBejoo/parking-space-monitor/internal/entity"
	"github.com/SangBejoo/parking-space-monitor/internal/repository"
	"github.com/xjem/t38c"
)

type MonitoringUsecase struct {
	TrxSupplyRepo repository.TrxSupplyRepository
	MapPlaceRepo  repository.MapPlaceRepository  // This is now using the interface
}

func NewMonitoringUsecase(trxRepo repository.TrxSupplyRepository, mapRepo repository.MapPlaceRepository) *MonitoringUsecase {
	return &MonitoringUsecase{
		TrxSupplyRepo: trxRepo,
		MapPlaceRepo:  mapRepo,
	}
}

func (u *MonitoringUsecase) GetMonitoringData() ([]entity.PlaceMonitoring, error) {
	host := os.Getenv("TILE38_HOST")
	port := os.Getenv("TILE38_PORT")
	if host == "" {
		host = "localhost"
	}
	if port == "" {
		port = "9851"
	}

	client, err := t38c.New(t38c.Config{
		Address: fmt.Sprintf("%s:%s", host, port),
	})
	if (err != nil) {
		return nil, fmt.Errorf("failed to connect to Tile38: %v", err)
	}
	defer client.Close()

	ctx := context.Background()

	// Retrieve data from repositories
	trxSupplies, err := u.TrxSupplyRepo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("failed to get supplies: %v", err)
	}
	mapPlaces, err := u.MapPlaceRepo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("failed to get places: %v", err)
	}

	// Clear existing data
	// Note: Tile38 does not support FlushDB, so this step is skipped.

	// Insert MapPlace polygons into Tile38
	for _, place := range mapPlaces {
		key := fmt.Sprintf("place:%d", place.ID)
		err := client.Keys.Set("places", key).
			String(place.Polygon).
			Do(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to set place %d: %v", place.ID, err)
		}
	}

	// Insert TrxSupply locations into Tile38
	for _, supply := range trxSupplies {
		err := client.Keys.Set("drivers", supply.DriverID).
			Point(supply.Latitude, supply.Longitude).
			Do(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to set driver %s: %v", supply.DriverID, err)
		}
	}

	// Perform spatial queries to match drivers with places
	var monitoringData []entity.PlaceMonitoring
	for _, place := range mapPlaces {
		drivers, err := u.TrxSupplyRepo.FindDriversInPlace(place.ID)
		if err != nil {
			return nil, fmt.Errorf("failed to query drivers in place %d: %v", place.ID, err)
		}

		monitoringData = append(monitoringData, entity.PlaceMonitoring{
			ID:      place.ID,
			Total:   len(drivers),
			Polygon: place.Polygon,
			Driver:  drivers,
		})
	}

	return monitoringData, nil
}