package blockchain 

import (
	"github.com/dougkirkley/helium-go/helium"
)

type Oracle struct {
	c *helium.Client
}

func NewOracle(c *helium.Client) *Oracle {
	return &Oracle{c}
}