package entity

import "time"

type MapPlace struct {
    ID         int       `json:"id"`
    HexagonID  string    `json:"hexagon_id"`
    PlaceID    string    `json:"place_id"`
    PlaceName  string    `json:"place_name"`
    Polygon    string    `json:"polygon"`
    CreatedAt  time.Time `json:"created_at"`
    UpdatedAt  time.Time `json:"updated_at"`
}