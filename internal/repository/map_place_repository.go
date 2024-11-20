package repository

import (
	"database/sql"
	"github.com/SangBejoo/parking-space-monitor/internal/entity"
)

type MapPlaceRepository interface {
	Create(place entity.MapPlace) error
	GetAll() ([]entity.MapPlace, error)
	FindByPoint(lat, lng float64) (*entity.MapPlace, error)
}

type mapPlaceRepositoryImpl struct {
	db *sql.DB
}

func NewMapPlaceRepository(db *sql.DB) MapPlaceRepository {
	return &mapPlaceRepositoryImpl{db: db}
}

func (r *mapPlaceRepositoryImpl) Create(place entity.MapPlace) error {
	query := `
		INSERT INTO map_places (hexagon_id, place_id, place_name, polygon)
		VALUES ($1, $2, $3, ST_GeomFromText($4, 4326))
		RETURNING id`
	
	return r.db.QueryRow(query, 
		place.HexagonID,
		place.PlaceID,
		place.PlaceName,
		place.Polygon).Scan(&place.ID)
}

func (r *mapPlaceRepositoryImpl) GetAll() ([]entity.MapPlace, error) {
	query := `
		SELECT id, hexagon_id, place_id, place_name, 
			   ST_AsText(polygon) as polygon,
			   created_at, updated_at 
		FROM map_places`
	
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var places []entity.MapPlace
	for rows.Next() {
		var mp entity.MapPlace
		err = rows.Scan(&mp.ID, &mp.HexagonID, &mp.PlaceID, &mp.PlaceName, &mp.Polygon, &mp.CreatedAt, &mp.UpdatedAt)
		if err != nil {
			return nil, err
		}
		places = append(places, mp)
	}
	return places, nil
}

func (r *mapPlaceRepositoryImpl) FindByPoint(lat, lng float64) (*entity.MapPlace, error) {
	query := `
		SELECT id, hexagon_id, place_id, place_name, 
			   ST_AsText(polygon) as polygon,
			   created_at, updated_at 
		FROM map_places
		WHERE ST_Contains(polygon, ST_SetSRID(ST_MakePoint($1, $2), 4326))
		LIMIT 1`
	
	var place entity.MapPlace
	err := r.db.QueryRow(query, lng, lat).Scan(
		&place.ID,
		&place.HexagonID,
		&place.PlaceID,
		&place.PlaceName,
		&place.Polygon,
		&place.CreatedAt,
		&place.UpdatedAt)
	
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &place, nil
}