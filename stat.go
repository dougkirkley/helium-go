package helium

import (
	"bytes"
	"net/http"
	"encoding/json"
)

// Stat handles api endpoint /stats docs located at https://docs.helium.com/api/blockchain/stats
type Stat struct {
	c *Client
}

// Stat returns the Stat client
func (c *Client) Stat() *Stat {
	return &Stat{c}
}

// Stats holds data for the /stats endpoint
type Stats struct {
	Data StatsData `json:"data"`
}

// BlockTimes times for the block
type BlockTimes struct {
	LastDay   LastDay   `json:"last_day"`
	LastHour  LastHour  `json:"last_hour"`
	LastMonth LastMonth `json:"last_month"`
	LastWeek  LastWeek  `json:"last_week"`
}

// ChallengeCounts count stats for challenges
type ChallengeCounts struct {
	Active  int `json:"active"`
	LastDay int `json:"last_day"`
}

// Counts struct
type Counts struct {
	Blocks          int `json:"blocks"`
	Challenges      int `json:"challenges"`
	Cities          int `json:"cities"`
	ConsensusGroups int `json:"consensus_groups"`
	Countries       int `json:"countries"`
	Hotspots        int `json:"hotspots"`
	Transactions    int `json:"transactions"`
}

// ElectionTimes time of elections
type ElectionTimes struct {
	LastDay   LastDay   `json:"last_day"`
	LastHour  LastHour  `json:"last_hour"`
	LastMonth LastMonth `json:"last_month"`
	LastWeek  LastWeek  `json:"last_week"`
}

// Fees struct
type Fees struct {
	LastDay   LastDay   `json:"last_day"`
	LastMonth LastMonth `json:"last_month"`
	LastWeek  LastWeek  `json:"last_week"`
}

// StateChannelCounts struct
type StateChannelCounts struct {
	LastDay   LastDay   `json:"last_day"`
	LastMonth LastMonth `json:"last_month"`
	LastWeek  LastWeek  `json:"last_week"`
}
type StatsData struct {
	BlockTimes         BlockTimes         `json:"block_times"`
	ChallengeCounts    ChallengeCounts    `json:"challenge_counts"`
	Counts             Counts             `json:"counts"`
	ElectionTimes      ElectionTimes      `json:"election_times"`
	Fees               Fees               `json:"fees"`
	StateChannelCounts StateChannelCounts `json:"state_channel_counts"`
	TokenSupply        float64            `json:"token_supply"`
}

type TokenSupply struct {
	Data TokenSupplyData `json:"data"`
}
type TokenSupplyData struct {
	TokenSupply float64 `json:"token_supply"`
}

/* 
List Retrieve basic stats for the blockchain such as total token supply, 
and average block and election times over a number of intervals.
*/
func (s *Stat) List() (*Stats, error) {
	resp, err := s.c.Request(http.MethodGet, "/stats", new(bytes.Buffer), nil)
	if err != nil {
		return &Stats{}, err
	}
	defer resp.Body.Close()
	
	var stats *Stats
	err = json.NewDecoder(resp.Body).Decode(&stats)
	if err != nil {
		return &Stats{}, err
	}
	return stats, nil
}

// TokenSupply Returns the circulating token supply
func (s *Stat) TokenSupply() (*TokenSupply, error) {
	resp, err := s.c.Request(http.MethodGet, "/stats/token_supply", new(bytes.Buffer), nil)
	if err != nil {
		return &TokenSupply{}, err
	}
	defer resp.Body.Close()

	var tokenSupply *TokenSupply
	err = json.NewDecoder(resp.Body).Decode(&tokenSupply)
	if err != nil {
		return &TokenSupply{}, err
	}
	return tokenSupply, nil
}