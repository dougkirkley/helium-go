package helium 

type Transaction struct {
	c *Client
}

func (c *Client) Transaction() *Transaction {
	return &Transaction{c}
}