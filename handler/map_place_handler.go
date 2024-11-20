
package handler

import (
	"encoding/json"
	"net/http"
	"github.com/SangBejoo/parking-space-monitor/entity"
	"github.com/SangBejoo/parking-space-monitor/usecase"
)

type MapPlaceHandler struct {
	usecase usecase.MapPlaceUsecase
}

func NewMapPlaceHandler(u usecase.MapPlaceUsecase) *MapPlaceHandler {
	return &MapPlaceHandler{
		usecase: u,
	}
}

func (h *MapPlaceHandler) CreateMapPlace(w http.ResponseWriter, r *http.Request) {
	var mapPlace entity.MapPlace
	if err := json.NewDecoder(r.Body).Decode(&mapPlace); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.usecase.Create(mapPlace); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(mapPlace)
}

func (h *MapPlaceHandler) GetAllMapPlace(w http.ResponseWriter, r *http.Request) {
	places, err := h.usecase.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(places)
}