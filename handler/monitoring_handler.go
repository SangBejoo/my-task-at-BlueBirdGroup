package handler

import (
	"encoding/json"
	"net/http"

	"github.com/SangBejoo/parking-space-monitor/usecase"
)

type MonitoringHandler struct {
	Usecase *usecase.MonitoringUsecase
}

func NewMonitoringHandler(u *usecase.MonitoringUsecase) *MonitoringHandler {
	return &MonitoringHandler{Usecase: u}
}

func (h *MonitoringHandler) GetMonitoring(w http.ResponseWriter, r *http.Request) {
	data, err := h.Usecase.GetMonitoringData()
	if err != nil {
		http.Error(w, "Unable to get monitoring data", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
