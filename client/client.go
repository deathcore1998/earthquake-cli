package client

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/deathcore1998/earthquake-cli/models"
)

const baseURL = "https://earthquake.usgs.gov/earthquakes/feed/v1.0/summary/4.5_day.geojson"

func FetchEarthquakes() (models.EarthquakeResponse, error) {
	resp, err := http.Get(baseURL)
	if err != nil {
		return models.EarthquakeResponse{}, fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return models.EarthquakeResponse{}, fmt.Errorf("unexpected status: %d", resp.StatusCode)
	}

	var result models.EarthquakeResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return models.EarthquakeResponse{}, fmt.Errorf("JSON decode failed: %w", err)
	}
	return result, nil
}
