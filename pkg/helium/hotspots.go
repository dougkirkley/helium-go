package helium

type Hotspot struct {
	c *Client
}

func (c *Client) Hotspot() *Hotspot {
	return &Hotspot{c}
}