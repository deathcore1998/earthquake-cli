package main

import (
	"fmt"
	"os"

	"github.com/deathcore1998/earthquake-cli/client"
)

func main() {
	earthquakes, err := client.FetchEarthquakes()
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	fmt.Printf("Total number earthquakes: %d\n", earthquakes.Metadata.Count)
	for _, value := range earthquakes.Features {
		fmt.Printf("M%.1f — %s\n", value.Properties.Mag, value.Properties.Place)
	}
}
