package blockchain 

import (
	helium "github.com/dougkirkley/helium-go/pkg/client"
)

type Transaction struct {
	c *helium.Client
}

func (c *helium.Client) Transaction() *Transaction {
	return &Transaction{c}
}