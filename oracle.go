package helium

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Oracle struct {
	c *Client
}

func (c *Client) Oracle() *Oracle {
	return &Oracle{c}
}

type OraclePrices struct {
	Data   []OraclePriceData `json:"data"`
	Cursor string            `json:"cursor"`
}

type OraclePrice struct {
	Data OraclePriceData `json:"data"`
}

type OraclePriceData struct {
	Price int `json:"price"`
	Block int `json:"block"`
}

type OraclePriceStats struct {
	Data OraclePriceStatsData `json:"data"`
	Meta Meta                 `json:"meta"`
}

type OraclePriceStatsData struct {
	Avg    float64 `json:"avg"`
	Max    float64 `json:"max"`
	Median float64 `json:"median"`
	Min    float64 `json:"min"`
	Stddev float64 `json:"stddev"`
}

type OraclePriceActivity struct {
	Cursor string `json:"cursor"`
	Data   []OraclePriceActivityData `json:"data"`
}

type OraclePriceActivityData struct {
	BlockHeight int    `json:"block_height"`
	Fee         int    `json:"fee"`
	Hash        string `json:"hash"`
	Height      int    `json:"height"`
	Price       int    `json:"price"`
	PublicKey   string `json:"public_key"`
	Time        int    `json:"time"`
	Type        string `json:"type"`
}

type OraclePriceListInput struct {
	Cursor string
}

type OraclePriceStatsInput struct {
	MinTime string
	MaxTime string
}

type OraclePriceBlockInput struct {
	ID string
}

type OraclePriceActivityInput struct {
	Cursor string
}

// List The current and historical Oracle Prices and at which block they took effect.
func (o *Oracle) List(input *OraclePriceListInput) (*OraclePrices, error) {
	params := make(map[string]string)
	params["cursor"] = input.Cursor
	resp, err := o.c.Request(http.MethodGet, "/oracle/prices", new(bytes.Buffer), params)
	if err != nil {
		return &OraclePrices{}, err
	}
	defer resp.Body.Close()

	var oraclePrices *OraclePrices
	err = json.NewDecoder(resp.Body).Decode(&oraclePrices)
	if err != nil {
		return &OraclePrices{}, err
	}
	return oraclePrices, nil
}

// Current The current Oracle Price and at which block it took effect.
func (o *Oracle) Current() (*OraclePrice, error) {
	resp, err := o.c.Request(http.MethodGet, "/oracle/prices/current", new(bytes.Buffer), nil)
	if err != nil {
		return &OraclePrice{}, err
	}
	defer resp.Body.Close()

	var oraclePrice *OraclePrice
	err = json.NewDecoder(resp.Body).Decode(&oraclePrice)
	if err != nil {
		return &OraclePrice{}, err
	}
	return oraclePrice, nil
}

// Stats Gets statistics on Oracle prices.
func (o *Oracle) Stats(input *OraclePriceStatsInput) (*OraclePriceStats, error) {
	params := make(map[string]string)
	params["min_time"] = input.MinTime
	params["max_time"] = input.MaxTime
	resp, err := o.c.Request(http.MethodGet, "/oracle/prices/stats", new(bytes.Buffer), params)
	if err != nil {
		return &OraclePriceStats{}, err
	}
	defer resp.Body.Close()

	var oraclePriceStats *OraclePriceStats
	err = json.NewDecoder(resp.Body).Decode(&oraclePriceStats)
	if err != nil {
		return &OraclePriceStats{}, err
	}
	return oraclePriceStats, nil
}

// Block Provides the oracle price at a specific block and at which block it initially took effect.
func (o *Oracle) Block(input *OraclePriceBlockInput) (*OraclePrice, error) {
	resp, err := o.c.Request(http.MethodGet, fmt.Sprintf("/oracle/prices/%s", input.ID), new(bytes.Buffer), nil)
	if err != nil {
		return &OraclePrice{}, err
	}
	defer resp.Body.Close()

	var oraclePrice *OraclePrice
	err = json.NewDecoder(resp.Body).Decode(&oraclePrice)
	if err != nil {
		return &OraclePrice{}, err
	}
	return oraclePrice, nil
}

// Activity List oracle price report transactions for all oracle keys.
func (o *Oracle) Activity(input *OraclePriceActivityInput) (*OraclePriceActivity, error) {
	params := make(map[string]string)
	params["cursor"] = input.Cursor
	resp, err := o.c.Request(http.MethodGet, "/oracle/prices/activity", new(bytes.Buffer), params)
	if err != nil {
		return &OraclePriceActivity{}, err
	}
	defer resp.Body.Close()

	var oraclePriceActivity *OraclePriceActivity
	err = json.NewDecoder(resp.Body).Decode(&oraclePriceActivity)
	if err != nil {
		return &OraclePriceActivity{}, err
	}
	return oraclePriceActivity, nil
}