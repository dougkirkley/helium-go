package blockchain

import (
	helium "github.com/dougkirkley/helium-go/pkg/client"
)

type Oracle struct {
	c *helium.Client
}

func (c *helium.Client) Oracle() *Oracle {
	return &Oracle{c}
}