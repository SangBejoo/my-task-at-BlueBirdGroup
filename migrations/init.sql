-- Enable PostGIS extension
CREATE EXTENSION IF NOT EXISTS postgis;

-- Create map_places table
CREATE TABLE IF NOT EXISTS map_places (
    id SERIAL PRIMARY KEY,
    hexagon_id VARCHAR(50),
    place_id VARCHAR(50),
    place_name VARCHAR(100),
    polygon geometry(POLYGON, 4326),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create trx_supply table
CREATE TABLE IF NOT EXISTS trx_supply (
    id SERIAL PRIMARY KEY,
    fleet_number VARCHAR(50),
    latitude DOUBLE PRECISION,
    longitude DOUBLE PRECISION,
    place_id INT,
    place_type_id INT,
    driver_id VARCHAR(50),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create indices for spatial queries
CREATE INDEX IF NOT EXISTS idx_map_places_polygon ON map_places USING GIST (polygon);

-- Insert test data
INSERT INTO map_places (hexagon_id, place_id, place_name, polygon) VALUES
('HEX001', 'PLACE001', 'Mall Taman Anggrek', 
 ST_GeomFromText('POLYGON((106.790123 -6.178901, 106.791234 -6.178901, 106.791234 -6.179012, 106.790123 -6.179012, 106.790123 -6.178901))', 4326));

-- Insert supply data that falls within the polygon
INSERT INTO trx_supply (fleet_number, latitude, longitude, driver_id, place_type_id) VALUES
('TAX-001', -6.178950, 106.790500, 'DRV001', 1);