package helium 

type Block struct {
	c *Client
}

func (c *Client) Block() *Block {
	return &Block{c}
}