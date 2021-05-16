package helium

type Oracle struct {
	c *Client
}

func (c *Client) Oracle() *Oracle {
	return &Oracle{c}
}

type OraclePrice struct {
	Data OraclePriceData `json:"data"`
}

type OraclePriceData struct {
	Price int `json:"price"`
	Block int `json:"block"`
}

