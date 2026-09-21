package models

type EarthquakeResponse struct {
	Type     string    `json:"type"`
	Metadata Metadata  `json:"metadata"`
	Features []Feature `json:"features"`
}

type Metadata struct {
	Generated int64  `json:"generated"`
	URL       string `json:"url"`
	Title     string `json:"title"`
	Status    int    `json:"status"`
	API       string `json:"api"`
	Count     int    `json:"count"`
}

type Feature struct {
	Type       string     `json:"type"`
	Properties Properties `json:"properties"`
	Geometry   Geometry   `json:"geometry"`
	ID         string     `json:"id"`
}

type Properties struct {
	Mag     float64  `json:"mag"`
	Place   string   `json:"place"`
	Time    int64    `json:"time"`
	Updated int64    `json:"updated"`
	Tz      *string  `json:"tz"`
	URL     string   `json:"url"`
	Detail  string   `json:"detail"`
	Felt    *int     `json:"felt"`
	Cdi     *float64 `json:"cdi"`
	Mmi     *float64 `json:"mmi"`
	Alert   *string  `json:"alert"`
	Status  string   `json:"status"`
	Tsunami int      `json:"tsunami"`
	Sig     int      `json:"sig"`
	Net     string   `json:"net"`
	Code    string   `json:"code"`
	IDs     string   `json:"ids"`
	Sources string   `json:"sources"`
	Types   string   `json:"types"`
	Nst     int      `json:"nst"`
	Dmin    float64  `json:"dmin"`
	Rms     float64  `json:"rms"`
	Gap     int      `json:"gap"`
	MagType string   `json:"magType"`
	Type    string   `json:"type"`
	Title   string   `json:"title"`
}

type Geometry struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}
