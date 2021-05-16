// +build ignore 
package main

import (
	"fmt"
	helium "github.com/dougkirkley/helium-go"
)

func main() {
	client := helium.DefaultClient()
	account, err := client.Account().Get("13buBykFQf5VaQtv7mWj2PBY9Lq4i1DeXhg7C4Vbu3ppzqqNkTH")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Account: %v", account.Data)
}