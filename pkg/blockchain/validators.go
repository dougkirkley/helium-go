package blockchain 

import (
	helium "github.com/dougkirkley/helium-go/pkg/client"
)

type Validator struct {
	c *helium.Client
}

func (c *helium.Client) Validator() *Validator {
	return &Validator{c}
}