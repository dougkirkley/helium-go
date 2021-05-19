// +build ignore 
package main

import (
	"fmt"
	helium "github.com/dougkirkley/helium-go"
)

func main() {
	client := helium.DefaultClient()
	height, err := client.Block().CurrentHeight(&helium.BlockCursorInput{})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Current Height: %v", height.Data.Height)
}