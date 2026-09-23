package main

import (
	"fmt"
	"os"

	"github.com/deathcore1998/earthquake-cli/client"
	"github.com/deathcore1998/earthquake-cli/store"
)

func main() {
	database, err := store.NewDB("earthquake.db")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	defer database.Close()

	earthquakes, err := client.FetchEarthquakes()
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	fmt.Printf("Total number earthquakes: %d\n", earthquakes.Metadata.Count)
	for _, value := range earthquakes.Features {
		if err := database.Save(value); err != nil {
			fmt.Println("Save error:", err)
			continue
		}
		fmt.Printf("M%.1f — %s\n", value.Properties.Mag, value.Properties.Place)
	}
}
