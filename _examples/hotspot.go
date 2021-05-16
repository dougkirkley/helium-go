// +build ignore 
package main

import (
	"fmt"
	helium "github.com/dougkirkley/helium-go"
)

func main() {
	client := helium.DefaultClient()
	hotspots, err := client.Hotspot().Distance(41.24450048207128, -73.93189556758152, 100)
	if err != nil {
		fmt.Println(err)
	}
	for _, hotspot := range hotspots.Data {
		fmt.Println(hotspot.Location)
	}
}