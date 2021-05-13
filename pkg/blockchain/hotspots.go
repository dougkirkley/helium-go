package blockchain 

import (
	helium "github.com/dougkirkley/helium-go/pkg/client"
)

type Hotspot struct {
	c *helium.Client
}

func (c *helium.Client) Hotspot() *Hotspot {
	return &Hotspot{c}
}