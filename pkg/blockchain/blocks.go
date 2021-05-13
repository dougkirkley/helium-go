package blockchain

import (
	helium "github.com/dougkirkley/helium-go/pkg/client"
)

type Block struct {
	c *helium.Client
}

func (c *helium.Client) Block() *Block {
	return &Block{c}
}