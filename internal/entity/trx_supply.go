package entity

import "time"

type TrxSupply struct {
    ID           int        `json:"id"`
    FleetNumber  string     `json:"fleet_number"`
    Latitude     float64    `json:"latitude"`
    Longitude    float64    `json:"longitude"`
    PlaceID      *int       `json:"place_id"`
    PlaceTypeID  *int       `json:"place_type_id"`
    DriverID     string     `json:"driver_id"`
    CreatedAt    time.Time  `json:"created_at"`
    UpdatedAt    time.Time  `json:"updated_at"`
}