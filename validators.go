package helium

type Validator struct {
	c *Client
}

func (c *Client) Validator() *Validator {
	return &Validator{c}
}