package handler

import (
	"encoding/json"
	"net/http"
	"github.com/SangBejoo/parking-space-monitor/internal/usecase"
	"github.com/SangBejoo/parking-space-monitor/internal/entity"
)

type TrxSupplyHandler struct {
	usecase usecase.TrxSupplyUsecase
}

func NewTrxSupplyHandler(u usecase.TrxSupplyUsecase) *TrxSupplyHandler {
	return &TrxSupplyHandler{
		usecase: u,
	}
}

func (h *TrxSupplyHandler) CreateTrxSupply(w http.ResponseWriter, r *http.Request) {
	var supply entity.TrxSupply
	if err := json.NewDecoder(r.Body).Decode(&supply); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate required fields
	if supply.FleetNumber == "" || supply.DriverID == "" {
		http.Error(w, "fleet_number and driver_id are required", http.StatusBadRequest)
		return
	}

	// Initialize null pointers if they weren't set in the JSON
	if supply.PlaceID == nil {
		supply.PlaceID = nil
	}
	if supply.PlaceTypeID == nil {
		supply.PlaceTypeID = nil
	}

	if err := h.usecase.Create(supply); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "success",
		"message": "Supply created successfully",
		"data": supply,
	})
}

func (h *TrxSupplyHandler) GetAllTrxSupply(w http.ResponseWriter, r *http.Request) {
	supplies, err := h.usecase.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "success",
		"data": supplies,
	})
}