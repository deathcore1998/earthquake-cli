package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/deathcore1998/earthquake-cli/client"
	"github.com/deathcore1998/earthquake-cli/store"
)

func main() {
	url := flag.String("url", "https://earthquake.usgs.gov/earthquakes/feed/v1.0/summary/4.5_day.geojson", "USGS earthquake feed URL")
	fetch := flag.Bool("fetch", false, "fetch earthquakes from USGS API and save to database")
	list := flag.Bool("list", false, "list earthquakes from database")
	limit := flag.Int("limit", 0, "maximum number of earthquakes to display (0 = all)")
	sortBy := flag.String("sort", "time", "sort by: time, mag, none")
	order := flag.String("order", "desc", "sort order: asc, desc")

	dbPath := flag.String("db", "earthquake.db", "path to SQLite database file")

	flag.Parse()

	if !*fetch && !*list {
		flag.Usage()
		os.Exit(1)
	}

	database, err := store.NewDB(*dbPath)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	defer database.Close()
	if *fetch {
		earthquakes, err := client.FetchEarthquakes(*url)
		if err != nil {
			fmt.Println("Error: ", err)
			os.Exit(1)
		}

		for _, future := range earthquakes.Features {
			err := database.Save(future)
			if err != nil {
				fmt.Println("Save error:", err)
			}
		}
	}

	if *list {
		futures, err := database.GetAll(*sortBy, *order)
		if err != nil {
			fmt.Println("Error: ", err)
			os.Exit(1)
		}

		count := len(futures)
		if *limit > 0 && *limit < count {
			count = *limit
		}

		for index := 0; index < count; index++ {
			fmt.Println(futures[index])
		}
	}
}
