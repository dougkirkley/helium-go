# helium-go

Go implementation of Helium API client

# Overview
The Helium API is a REST API service as defined by the blockhain-http service. This library supports the conventions as exposed by the API. This includes:

* Modular access to each of the main areas of the Helium API
* Support for paged responses

Contributions and helpful suggestions are always welcome

# Quickstart Instructions

Install the library
```go
go get github.com/dougkirkley/helium-go
```

# Example
Create a client to the default api.helium.io endpoint and ask for a given account.

```go
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
```
See the _examples folder and unit tests for more examples.