package helium

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type PendingTransaction struct {
	c *Client
}

func (c *Client) PendingTransaction() *PendingTransaction {
	return &PendingTransaction{c}
}

type PendingTransactions struct {
	Data []PendingTransactionData `json:"data"`
}

type PendingTransactionData struct {
	UpdatedAt    time.Time `json:"updated_at"`
	Type         string    `json:"type"`
	Txn          Txn       `json:"txn"`
	Status       string    `json:"status"`
	Hash         string    `json:"hash"`
	FailedReason string    `json:"failed_reason"`
	CreatedAt    time.Time `json:"created_at"`
}

type SubmittedHash struct {
	Data SubmittedHashData `json:"data"`
}

type SubmittedHashData struct {
	Hash string `json:"hash"`
}

type TransactionSubmitBody struct {
	Txn string `json:"txn"`
}

type PendingTransactionInput struct {
	ID string
}

type TransactionSubmitInput struct {
	Transaction string
}

// Get Fetches the status for a given pending transaction hash.
func (t *PendingTransaction) Get(input *PendingTransactionInput) (*PendingTransactions, error) {
	resp, err := t.c.Request(http.MethodGet, fmt.Sprintf("/pending_transactions/%s", input.ID), new(bytes.Buffer), nil)
	if err != nil {
		return &PendingTransactions{}, err
	}
	defer resp.Body.Close()

	var pendingTransactions *PendingTransactions
	err = json.NewDecoder(resp.Body).Decode(&pendingTransactions)
	if err != nil {
		return &PendingTransactions{}, err
	}
	return pendingTransactions, nil
}

// Submit New transactions can be submitted to the blockchain by sending a pending transaction.
func (t *PendingTransaction) Submit(input *TransactionSubmitInput) error {
	encodedTransaction := base64.StdEncoding.EncodeToString([]byte(input.Transaction))
	transactionData := TransactionSubmitBody{
		Txn: encodedTransaction,
	}
	body, err := json.Marshal(transactionData)
	if err != nil {
		return err
	}
	_, err = t.c.Request(http.MethodPost, "/pending_transactions", bytes.NewBuffer(body), nil)
	if err != nil {
		return err
	}
	return nil
}
