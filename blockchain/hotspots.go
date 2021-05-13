package blockchain

import (
	"github.com/dougkirkley/helium-go/helium"
)

type Hotspot struct {
	c *helium.Client
}

func NewHotspot(c *helium.Client) *Hotspot {
	return &Hotspot{c}
}