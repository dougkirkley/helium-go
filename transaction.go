package helium

import (
	"bytes"
	"fmt"
	"net/http"
	"encoding/json"
)

type Transaction struct {
	c *Client
}

func (c *Client) Transaction() *Transaction {
	return &Transaction{c}
}

type TransactionInfo struct {
	Data TransactionData `json:"data"`
}

// Get Fetch the transaction for a given hash.
func (t *Transaction) Get(hash string) (*TransactionInfo, error) {
	resp, err := t.c.Request(http.MethodGet, fmt.Sprintf("/transactions/%s", hash), new(bytes.Buffer), nil)
	if err != nil {
		return &TransactionInfo{}, err
	}
	defer resp.Body.Close()

	var transactionInfo *TransactionInfo
	err = json.NewDecoder(resp.Body).Decode(&transactionInfo)
	if err != nil {
		return &TransactionInfo{}, err
	}
	return transactionInfo, nil
}