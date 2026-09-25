package client

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/deathcore1998/earthquake-cli/models"
)

func FetchEarthquakes(url string) (models.EarthquakeResponse, error) {
	resp, err := http.Get(url)
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
