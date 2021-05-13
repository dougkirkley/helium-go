package blockchain

import (
	"github.com/dougkirkley/helium-go/helium"
)

type Transaction struct {
	c *helium.Client
}

func NewTransaction(c *helium.Client) *Transaction {
	return &Transaction{c}
}