package entity

type PlaceMonitoring struct {
    ID      int      `json:"id"`
    Total   int      `json:"total"`
    Polygon string   `json:"polygon"`
    Driver  []string `json:"driver"`
}