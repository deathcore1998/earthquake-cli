package store

import (
	"database/sql"
	"fmt"

	"github.com/deathcore1998/earthquake-cli/models"
	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	db *sql.DB
}

func (database *DB) createTable() error {
	query := `CREATE TABLE IF NOT EXISTS earthquakes (
		id          TEXT PRIMARY KEY,
		mag         REAL,
		place       TEXT,
		time        INTEGER,
		tsunami     INTEGER,
		url         TEXT,
		status      TEXT,
		longitude   REAL,
		latitude    REAL,
		depth       REAL	
	)`

	_, err := database.db.Exec(query)
	return err
}

func (database *DB) Close() error {
	return database.db.Close()
}

func (database *DB) Save(future models.Feature) error {
	query := `INSERT OR REPLACE INTO earthquakes 
	(id, mag, place, time, tsunami, url, status, longitude, latitude, depth)
	VALUES
	(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	_, err := database.db.Exec(
		query,
		future.ID,
		future.Properties.Mag,
		future.Properties.Place,
		future.Properties.Time,
		future.Properties.Tsunami,
		future.Properties.URL,
		future.Properties.Status,
		future.Longitude(),
		future.Latitude(),
		future.Depth())

	return err
}

func (database *DB) GetAll(sortBy, order string) ([]models.Feature, error) {
	query := `
		SELECT id, mag, place, time, tsunami, url, status, longitude, latitude, depth
		FROM earthquakes`

	validSort := map[string]bool{
		"time": true,
		"mag":  true,
	}

	if sortBy != "none" && validSort[sortBy] {
		dir := "DESC"
		if order == "asc" {
			dir = "ASC"
		}
		query += fmt.Sprintf(" ORDER BY %s %s", sortBy, dir)
	}

	rows, err := database.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("query: %w", err)
	}

	defer rows.Close()
	var features []models.Feature

	for rows.Next() {
		var feature models.Feature
		feature.Geometry.Coordinates = make([]float64, 3)

		err := rows.Scan(
			&feature.ID,
			&feature.Properties.Mag,
			&feature.Properties.Place,
			&feature.Properties.Time,
			&feature.Properties.Tsunami,
			&feature.Properties.URL,
			&feature.Properties.Status,
			&feature.Geometry.Coordinates[0],
			&feature.Geometry.Coordinates[1],
			&feature.Geometry.Coordinates[2],
		)

		if err != nil {
			return nil, fmt.Errorf("scan: %w", err)
		}
		features = append(features, feature)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows: %w", err)
	}

	return features, nil
}

func NewDB(path string) (*DB, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, fmt.Errorf("open db: %w", err)
	}

	// connection check
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("ping db: %w", err)
	}

	database := &DB{db: db}
	if err := database.createTable(); err != nil {
		return nil, fmt.Errorf("create table: %w", err)
	}

	return database, nil
}
