package helium

type Oracle struct {
	c *Client
}

func (c *Client) Oracle() *Oracle {
	return &Oracle{c}
}