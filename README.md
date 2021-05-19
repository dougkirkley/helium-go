# helium-go

Go implementation of Helium API client to access the public [Helium](https://helium.com) blockchain REST API.

# Overview
The Helium API is a REST API service as defined by the
[blockhain-http](https://github.com/helium/blockchain-http) service This library supports the conventions as exposed by the API. This includes:

* Modular access to each of the main areas of the Helium API
* Support for paged responses

Contributions and helpful suggestions are [always
welcome](https://github.com/dougkirkley/helium-go/issues)


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
	accounts, err := client.Account().List(&helium.AccountListInput{})
	if err != nil {
		fmt.Println(err)
	}
	for _, account := range accounts.Data {
		fmt.Println(account.Address)
	}
}
```
See the _examples folder and unit tests for more examples.