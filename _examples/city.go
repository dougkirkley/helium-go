// +build ignore 
package main

import (
	"fmt"
	helium "github.com/dougkirkley/helium-go"
)

func main() {
	client := helium.DefaultClient()
	input := &helium.CityInput{
		ID: "aG91c3RvbnRleGFzdW5pdGVkIHN0YXRlcw",
	}
	hotspots, err := client.City().Hotspots(input)
	if err != nil {
		fmt.Println(err)
	}
	for _, hotspot := range hotspots.Data {
		fmt.Println(hotspot.Name)
	}
}