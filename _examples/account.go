// +build ignore 
package main

import (
	"fmt"
	helium "github.com/dougkirkley/helium-go"
)

func main() {
	client := helium.DefaultClient()
	accounts, err := client.Account().List(helium.NoQuery)
	if err != nil {
		fmt.Println(err)
	}
	for _, account := range accounts.Data {
		fmt.Println(account.Address)
	}
}