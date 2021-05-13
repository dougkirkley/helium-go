package blockchain 

import (
	"github.com/dougkirkley/helium-go/helium"
)

type Validator struct {
	c *helium.Client
}

func NewValidator(c *helium.Client) *Validator {
	return &Validator{c}
}