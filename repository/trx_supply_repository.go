package repository

import (
    "database/sql"
    "github.com/SangBejoo/parking-space-monitor/entity"
)

type TrxSupplyRepository interface {
    Create(supply entity.TrxSupply) error
    GetAll() ([]entity.TrxSupply, error)
    FindDriversInPlace(placeID int) ([]string, error)
}

type trxSupplyRepositoryImpl struct {
    db *sql.DB
}

func NewTrxSupplyRepository(db *sql.DB) TrxSupplyRepository {
    return &trxSupplyRepositoryImpl{db: db}
}

func (r *trxSupplyRepositoryImpl) Create(supply entity.TrxSupply) error {
    query := `
        INSERT INTO trx_supply (fleet_number, latitude, longitude, driver_id, place_id, place_type_id)
        VALUES ($1, $2, $3, $4, $5, $6)
        RETURNING id`
    
    var placeIDSQL sql.NullInt64
    if supply.PlaceID != nil {
        placeIDSQL = sql.NullInt64{Int64: int64(*supply.PlaceID), Valid: true}
    }
    
    var placeTypeIDSQL sql.NullInt64
    if supply.PlaceTypeID != nil {
        placeTypeIDSQL = sql.NullInt64{Int64: int64(*supply.PlaceTypeID), Valid: true}
    }
    
    result := r.db.QueryRow(query, 
        supply.FleetNumber, 
        supply.Latitude, 
        supply.Longitude, 
        supply.DriverID,
        placeIDSQL,
        placeTypeIDSQL)
        
    return result.Scan(&supply.ID)
}

func (r *trxSupplyRepositoryImpl) GetAll() ([]entity.TrxSupply, error) {
    query := `
        SELECT id, fleet_number, latitude, longitude, driver_id,
               place_id, place_type_id, created_at, updated_at
        FROM trx_supply ORDER BY created_at DESC
    `
    
    rows, err := r.db.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var supplies []entity.TrxSupply
    for rows.Next() {
        var s entity.TrxSupply
        if err := rows.Scan(
            &s.ID,
            &s.FleetNumber,
            &s.Latitude,
            &s.Longitude,
            &s.DriverID,
            &s.PlaceID,     // Scanning into *int
            &s.PlaceTypeID, // Scanning into *int
            &s.CreatedAt,
            &s.UpdatedAt,
        ); err != nil {
            return nil, err
        }
        supplies = append(supplies, s)
    }
    return supplies, nil
}

func (r *trxSupplyRepositoryImpl) FindDriversInPlace(placeID int) ([]string, error) {
    query := `
        SELECT ts.driver_id 
        FROM trx_supply ts, map_places mp 
        WHERE mp.id = $1 
        AND ST_Contains(mp.polygon, ST_SetSRID(ST_MakePoint(ts.longitude, ts.latitude), 4326))
    `
    
    rows, err := r.db.Query(query, placeID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var drivers []string
    for rows.Next() {
        var driverID string
        if err := rows.Scan(&driverID); err != nil {
            return nil, err
        }
        drivers = append(drivers, driverID)
    }
    return drivers, nil
}