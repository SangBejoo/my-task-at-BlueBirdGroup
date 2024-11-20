package server

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/SangBejoo/parking-space-monitor/handler"
)

func StartServer(trxHandler *handler.TrxSupplyHandler, mapHandler *handler.MapPlaceHandler, monitoringHandler *handler.MonitoringHandler) {
	router := mux.NewRouter()

	// API routes
	api := router.PathPrefix("/api").Subrouter()

	// Places routes
	api.HandleFunc("/places", mapHandler.CreateMapPlace).Methods("POST")
	api.HandleFunc("/places", mapHandler.GetAllMapPlace).Methods("GET")

	// Supplies routes
	api.HandleFunc("/supplies", trxHandler.CreateTrxSupply).Methods("POST")
	api.HandleFunc("/supplies", trxHandler.GetAllTrxSupply).Methods("GET")

	// Monitoring routes
	api.HandleFunc("/monitoring", monitoringHandler.GetMonitoring).Methods("GET")

	log.Println("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}