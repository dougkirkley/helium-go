package helium

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Block handles api endpoint /accounts docs located at https://docs.helium.com/api/blockchain/blocks
type Block struct {
	c *Client
}

// Block returns the Block client
func (c *Client) Block() *Block {
	return &Block{c}
}

type Height struct {
	Data HeightData `json:"data"`
}
type HeightData struct {
	Height int `json:"height"`
}

type BlockStats struct {
	Data BlockStatsData `json:"data"`
}

type LastHour struct {
	Avg    float64 `json:"avg"`
	Stddev float64 `json:"stddev"`
}

type BlockStatsData struct {
	LastDay   LastDay   `json:"last_day"`
	LastHour  LastHour  `json:"last_hour"`
	LastMonth LastMonth `json:"last_month"`
	LastWeek  LastWeek  `json:"last_week"`
}

type Blocks struct {
	Data   []BlockData `json:"data"`
	Cursor string `json:"cursor"`
}
type BlockData struct {
	TransactionCount int    `json:"transaction_count"`
	Time             int    `json:"time"`
	SnapshotHash     string `json:"snapshot_hash"`
	PrevHash         string `json:"prev_hash"`
	Height           int    `json:"height"`
	Hash             string `json:"hash"`
}

type BlockHeight struct {
	Data BlockData `json:"data"`
}

type Transactions struct {
	Data []TransactionData `json:"data"`
}

type Hash struct {
	Data BlockData `json:"data"`
}

type TransactionData struct {
	Version         int     `json:"version,omitempty"`
	Type            string  `json:"type"`
	Time            int     `json:"time"`
	Signature       string  `json:"signature"`
	SecretHash      string  `json:"secret_hash,omitempty"`
	Owner           string  `json:"owner,omitempty"`
	OnionKeyHash    string  `json:"onion_key_hash"`
	Location        string  `json:"location,omitempty"`
	Lng             float64 `json:"lng,omitempty"`
	Lat             float64 `json:"lat,omitempty"`
	Height          int     `json:"height"`
	Hash            string  `json:"hash"`
	Fee             int     `json:"fee"`
	Challenger      string  `json:"challenger"`
	BlockHash       string  `json:"block_hash,omitempty"`
	Secret          string  `json:"secret,omitempty"`
	Path            []Path  `json:"path,omitempty"`
	ChallengerOwner string  `json:"challenger_owner,omitempty"`
	ChallengerLoc   string  `json:"challenger_loc,omitempty"`
}

type HashTransactions struct {
	Data []TransactionData `json:"data"`
}

// List Retrieves block descriptions.
func (b *Block) List(cursor string) (*Blocks, error) {
	resp, err := b.c.Request(http.MethodGet, "/blocks", nil)
	if err != nil {
		return &Blocks{}, err
	}
	var blocks *Blocks
	err = json.Unmarshal(resp, &blocks)
	if err != nil {
		return &Blocks{}, err
	}
	return blocks, nil
}

// Get Get block descriptor for block at height
func (b *Block) Get(hash string) (*Block, error) {
	resp, err := b.c.Request(http.MethodGet, fmt.Sprintf("/blocks/%s", hash), nil)
	if err != nil {
		return &Block{}, err
	}
	var block *Block
	err = json.Unmarshal(resp, &block)
	if err != nil {
		return &Block{}, err
	}
	return block, nil
}

// CurrentHeight Gets the current height of the blockchainn.
func (b *Block) CurrentHeight(cursor string) (*Height, error) {
	resp, err := b.c.Request(http.MethodGet, "/blocks/height", nil)
	if err != nil {
		return &Height{}, err
	}
	var height *Height
	err = json.Unmarshal(resp, &height)
	if err != nil {
		return &Height{}, err
	}
	return height, nil
}

// Stats Get statistics on block production times.
func (b *Block) Stats(cursor string) (*BlockStats, error) {
	resp, err := b.c.Request(http.MethodGet, "/blocks/stats", nil)
	if err != nil {
		return &BlockStats{}, err
	}
	var stats *BlockStats
	err = json.Unmarshal(resp, &stats)
	if err != nil {
		return &BlockStats{}, err
	}
	return stats, nil
}

// GetHeight Get block descriptor for block at height
func (b *Block) GetHeight(height string) (*BlockHeight, error) {
	resp, err := b.c.Request(http.MethodGet, fmt.Sprintf("/blocks/%s/height", height), nil)
	if err != nil {
		return &BlockHeight{}, err
	}
	var block *BlockHeight
	err = json.Unmarshal(resp, &block)
	if err != nil {
		return &BlockHeight{}, err
	}
	return block, nil
}

// Transactions Get transactions for a block at a given height.
func (b *Block) Transactions(height string) (*Transactions, error) {
	resp, err := b.c.Request(http.MethodGet, fmt.Sprintf("/blocks/%s/transactions", height), nil)
	if err != nil {
		return &Transactions{}, err
	}
	var transactions *Transactions
	err = json.Unmarshal(resp, &transactions)
	if err != nil {
		return &Transactions{}, err
	}
	return transactions, nil
}