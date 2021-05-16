// +build ignore 
package main

import (
	"fmt"
	helium "github.com/dougkirkley/helium-go"
)

func main() {
	client := helium.DefaultClient()
	hotspots, err := client.City().Hotspots("aG91c3RvbnRleGFzdW5pdGVkIHN0YXRlcw")
	if err != nil {
		fmt.Println(err)
	}
	for _, hotspot := range hotspots.Data {
		fmt.Println(hotspot.Name)
	}
}