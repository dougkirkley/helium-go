package blockchain 

import (
	"github.com/dougkirkley/helium-go/helium"
)

type Block struct {
	c *helium.Client
}

func NewBlock(c *helium.Client) *Block {
	return &Block{c}
}