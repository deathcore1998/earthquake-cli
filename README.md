# Earthquake CLI

A simple CLI tool for fetching and storing earthquake data from the USGS API. Written in Go.

## Features
```
- Fetches earthquake data from USGS API
- Stores data in SQLite database
- Lists earthquakes from database with sorting and filtering
- Colored output based on magnitude
- Progress bar during fetch
```

## How to run

### Build

go build -o earthquake-cli

### Fetch and save earthquakes

./earthquake-cli --fetch

### List earthquakes from database

./earthquake-cli --list

### List with limit and sorting

./earthquake-cli --list --limit=5 --sort=mag --order=desc

## Flags
```
| Flag      | Default           | Description                                           |
|-----------|-------------------|-------------------------------------------------------|
| `--fetch` | `false`           | Fetch earthquakes from USGS API and save to database  |
| `--list`  | `false`           | List earthquakes from database                        |
| `--limit` | `0`               | Maximum number of earthquakes to display (0 = all)    |
| `--sort`  | `time`            | Sort by: `time`, `mag`, `none`                        |
| `--order` | `desc`            | Sort order: `asc`, `desc`                             |
| `--db`    | `earthquake.db`   | Path to SQLite database file                          |
| `--url`   | USGS feed URL     | USGS earthquake feed URL                              |
```
## Example output
```
go run main.go --list --limit=10 --sort=mag --order=desc

M6.4 (2026-09-20 12:17) — 49 km NNE of Kainantu, Papua New Guinea
M5.9 (2026-10-04 01:55) — 39 km WSW of Tambolaka, Indonesia
M5.7 (2026-09-23 17:41) — 180 km NW of Hihifo, Tonga
M5.7 (2026-09-20 23:54) — south of Africa
M5.5 (2026-09-21 04:41) — 35 km NNE of Ruteng, Indonesia
M5.5 (2026-09-19 02:18) — Kermadec Islands region
M5.5 (2026-10-04 14:12) — Volcano Islands, Japan region
M5.4 (2026-09-20 15:45) — 68 km NNE of Kainantu, Papua New Guinea
M5.3 (2026-09-25 14:12) — Vanuatu region
M5.3 (2026-09-25 02:30) — 72 km NW of Finschhafen, Papua New Guinea
```
## Project structure
```
earthquake-cli/
├── main.go              # entry point, CLI flags
├── client/
│   └── client.go        # USGS API client
├── models/
│   └── models.go        # data structures
├── store/
    ├── store.go         # SQLite storage
    └── store_test.go    # tests
```
## Tests

go test ./...

## Tech stack

- Go 1.27.1
- SQLite (via `github.com/mattn/go-sqlite3`)
- Progress bar (via `github.com/schollz/progressbar/v3`)
