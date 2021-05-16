package helium

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Account handles api endpoint /accounts docs located at https://docs.helium.com/api/blockchain/accounts
type Account struct {
	c *Client
}

// Account returns the Account client
func (c *Client) Account() *Account {
	return &Account{c}
}

type Accounts struct {
	Data   []AccountData `json:"data"`
	Cursor string        `json:"cursor"`
}

type UserAccount struct {
	Data AccountData `json:"data"`
}

type AccountData struct {
	SecNonce   int    `json:"sec_nonce"`
	SecBalance int    `json:"sec_balance"`
	Nonce      int    `json:"nonce"`
	DcNonce    int    `json:"dc_nonce"`
	DcBalance  int    `json:"dc_balance"`
	Block      int    `json:"block"`
	Balance    int    `json:"balance"`
	Address    string `json:"address"`
}

type Status struct {
	Online      string   `json:"online"`
	ListenAddrs []string `json:"listen_addrs"`
	Height      int      `json:"height"`
}

type Geocode struct {
	ShortStreet  string `json:"short_street"`
	ShortState   string `json:"short_state"`
	ShortCountry string `json:"short_country"`
	ShortCity    string `json:"short_city"`
	LongStreet   string `json:"long_street"`
	LongState    string `json:"long_state"`
	LongCountry  string `json:"long_country"`
	LongCity     string `json:"long_city"`
	CityID       string `json:"city_id"`
}

type Ouis struct {
	Data   []OuiData `json:"data"`
	Cursor string    `json:"cursor"`
}

type Subnets struct {
	Mask int `json:"mask"`
	Base int `json:"base"`
}

type OuiData struct {
	Subnets   []Subnets `json:"subnets"`
	Owner     string    `json:"owner"`
	Oui       int       `json:"oui"`
	Nonce     int       `json:"nonce"`
	Block     int       `json:"block"`
	Addresses []string  `json:"addresses"`
}

type Activity struct {
	Data   []ActivityData `json:"data"`
	Cursor string         `json:"cursor"`
}

type Reward struct {
	Type    string `json:"type"`
	Gateway string `json:"gateway"`
	Amount  int    `json:"amount"`
	Account string `json:"account"`
}

type ActivityData struct {
	Type       string   `json:"type"`
	Time       int      `json:"time"`
	StartEpoch int      `json:"start_epoch"`
	Rewards    []Reward `json:"rewards"`
	Height     int      `json:"height"`
	Hash       string   `json:"hash"`
	EndEpoch   int      `json:"end_epoch"`
}

type ActivityCount struct {
	Data CountsData `json:"data"`
}

type CountsData struct {
	AddGatewayV1     int `json:"add_gateway_v1"`
	AssertLocationV1 int `json:"assert_location_v1"`
}

type Elections struct {
	Data   []ElectionData `json:"data"`
	Cursor string         `json:"cursor"`
}

type ElectionData struct {
	Type    string   `json:"type"`
	Time    int      `json:"time"`
	Proof   string   `json:"proof"`
	Members []string `json:"members"`
	Height  int      `json:"height"`
	Hash    string   `json:"hash"`
	Delay   int      `json:"delay"`
}

type Challenges struct {
	Data   []ChallengeData `json:"data"`
	Cursor string          `json:"cursor"`
}

type Witness struct {
	Timestamp  int64   `json:"timestamp"`
	Snr        int     `json:"snr"`
	Signal     int     `json:"signal"`
	PacketHash string  `json:"packet_hash"`
	Owner      string  `json:"owner"`
	Location   string  `json:"location"`
	IsValid    bool    `json:"is_valid"`
	Gateway    string  `json:"gateway"`
	Frequency  float64 `json:"frequency"`
	Datarate   string  `json:"datarate"`
	Channel    int     `json:"channel"`
}

type Receipt struct {
	Timestamp int64       `json:"timestamp"`
	Snr       int         `json:"snr"`
	Signal    int         `json:"signal"`
	Origin    string      `json:"origin"`
	Gateway   string      `json:"gateway"`
	Frequency int         `json:"frequency"`
	Datarate  interface{} `json:"datarate"`
	Data      string      `json:"data"`
	Channel   int         `json:"channel"`
}

type Path struct {
	Witnesses          []Witness `json:"witnesses"`
	Receipt            Receipt   `json:"receipt"`
	Geocode            Geocode   `json:"geocode"`
	ChallengeeOwner    string    `json:"challengee_owner"`
	ChallengeeLon      float64   `json:"challengee_lon"`
	ChallengeeLocation string    `json:"challengee_location"`
	ChallengeeLat      float64   `json:"challengee_lat"`
	Challengee         string    `json:"challengee"`
}

type ChallengeData struct {
	Type               string  `json:"type"`
	Time               int     `json:"time"`
	Secret             string  `json:"secret"`
	RequestBlockHash   string  `json:"request_block_hash"`
	Path               []Path  `json:"path"`
	OnionKeyHash       string  `json:"onion_key_hash"`
	Height             int     `json:"height"`
	Hash               string  `json:"hash"`
	Fee                int     `json:"fee"`
	ChallengerOwner    string  `json:"challenger_owner"`
	ChallengerLon      float64 `json:"challenger_lon"`
	ChallengerLocation string  `json:"challenger_location"`
	ChallengerLat      float64 `json:"challenger_lat"`
	Challenger         string  `json:"challenger"`
}

type AccountPendingTransactions struct {
	Data   []AccountPendingTransactionData `json:"data"`
	Cursor string                          `json:"cursor"`
}

type Payments struct {
	Amount int    `json:"amount"`
	Payee  string `json:"payee"`
}

type Txn struct {
	Fee       int        `json:"fee"`
	Nonce     int        `json:"nonce"`
	Payer     string     `json:"payer"`
	Payments  []Payments `json:"payments"`
	Signature string     `json:"signature"`
}

type AccountPendingTransactionData struct {
	CreatedAt    time.Time `json:"created_at"`
	FailedReason string    `json:"failed_reason"`
	Hash         string    `json:"hash"`
	Status       string    `json:"status"`
	Txn          Txn       `json:"txn"`
	Type         string    `json:"type"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type Rewards struct {
	Data []RewardData `json:"data"`
}

type RewardData struct {
	CreatedAt    time.Time `json:"created_at"`
	FailedReason string    `json:"failed_reason"`
	Hash         string    `json:"hash"`
	Status       string    `json:"status"`
	Txn          Txn       `json:"txn"`
	Type         string    `json:"type"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type RewardSum struct {
	Data RewardSumData `json:"data"`
}

type RewardSumData struct {
	MaxTime time.Time `json:"max_time"`
	MinTime time.Time `json:"min_time"`
	Sum     string    `json:"sum"`
}

type AccountStats struct {
	Data AccountStatsData `json:"data"`
}

type LastWeek struct {
	Timestamp time.Time `json:"timestamp"`
	Balance   int64     `json:"balance"`
}

type LastMonth struct {
	Timestamp time.Time `json:"timestamp"`
	Balance   int64     `json:"balance"`
}

type LastDay struct {
	Timestamp time.Time `json:"timestamp"`
	Balance   int64     `json:"balance"`
}

type AccountStatsData struct {
	LastWeek  []LastWeek  `json:"last_week"`
	LastMonth []LastMonth `json:"last_month"`
	LastDay   []LastDay   `json:"last_day"`
}

// List Retrieves the current set of known accounts
func (a *Account) List(cursor string) (*Accounts, error) {
	params := make(map[string]string)
	if len(cursor) > 0 {
		params["cursor"] = cursor
	}
	resp, err := a.c.Request(http.MethodGet, "/accounts", new(bytes.Buffer), params)
	if err != nil {
		return &Accounts{}, err
	}
	defer resp.Body.Close()

	var accounts *Accounts
	err = json.NewDecoder(resp.Body).Decode(&accounts)
	if err != nil {
		return &Accounts{}, err
	}
	return accounts, nil
}

// Richest Returns up to 100 of the accounts sorted by highest token balance.
func (a *Account) Richest(limit int) (*Accounts, error) {
	params := make(map[string]string)
	if limit > 0 {
		params["limit"] = fmt.Sprintf("%v", limit)
	}
	resp, err := a.c.Request(http.MethodGet, "/accounts/rich", new(bytes.Buffer), params)
	if err != nil {
		return &Accounts{}, err
	}
	defer resp.Body.Close()

	var accounts *Accounts
	err = json.NewDecoder(resp.Body).Decode(&accounts)
	if err != nil {
		return &Accounts{}, err
	}
	return accounts, nil
}

// Get Retrieve a specific account record.
func (a *Account) Get(accountID string) (*UserAccount, error) {
	resp, err := a.c.Request(http.MethodGet, fmt.Sprintf("/accounts/%s", accountID), new(bytes.Buffer), nil)
	if err != nil {
		return &UserAccount{}, err
	}
	defer resp.Body.Close()

	var account *UserAccount
	err = json.NewDecoder(resp.Body).Decode(&account)
	if err != nil {
		return &UserAccount{}, err
	}
	return account, nil
}

// Hotspots Fetches hotspots owned by a given account address.
func (a *Account) Hotspots(accountID string) (*Hotspots, error) {
	resp, err := a.c.Request(http.MethodGet, fmt.Sprintf("/accounts/%s/hotspots", accountID), new(bytes.Buffer), nil)
	if err != nil {
		return &Hotspots{}, err
	}
	defer resp.Body.Close()

	var hotspots *Hotspots
	err = json.NewDecoder(resp.Body).Decode(&hotspots)
	if err != nil {
		return &Hotspots{}, err
	}
	return hotspots, nil
}

// Ouis Fetches OUIs owned by a given account address.
func (a *Account) Ouis(accountID string) (*Ouis, error) {
	resp, err := a.c.Request(http.MethodGet, fmt.Sprintf("/accounts/%s/ouis", accountID), new(bytes.Buffer), nil)
	if err != nil {
		return &Ouis{}, err
	}
	defer resp.Body.Close()

	var ouis *Ouis
	err = json.NewDecoder(resp.Body).Decode(&ouis)
	if err != nil {
		return &Ouis{}, err
	}
	return ouis, nil
}

// Activity Fetches transactions that indicate activity for an account.
func (a *Account) Activity(accountID string) (*Activity, error) {
	resp, err := a.c.Request(http.MethodGet, fmt.Sprintf("/accounts/%s/activity", accountID), new(bytes.Buffer), nil)
	if err != nil {
		return &Activity{}, err
	}
	defer resp.Body.Close()

	var activity *Activity
	err = json.NewDecoder(resp.Body).Decode(&activity)
	if err != nil {
		return &Activity{}, err
	}
	return activity, nil
}

// ActivityCount Count transactions that indicate activity for an account.
func (a *Account) ActivityCount(accountID string) (*ActivityCount, error) {
	resp, err := a.c.Request(http.MethodGet, fmt.Sprintf("/accounts/%s/activity/count", accountID), new(bytes.Buffer), nil)
	if err != nil {
		return &ActivityCount{}, err
	}
	defer resp.Body.Close()

	var activityCount *ActivityCount
	err = json.NewDecoder(resp.Body).Decode(&activityCount)
	if err != nil {
		return &ActivityCount{}, err
	}
	return activityCount, nil
}

// Elections Fetches elections that hotspots for the given account are elected in.
func (a *Account) Elections(accountID string) (*Elections, error) {
	resp, err := a.c.Request(http.MethodGet, fmt.Sprintf("/accounts/%s/elections", accountID), new(bytes.Buffer), nil)
	if err != nil {
		return &Elections{}, err
	}
	defer resp.Body.Close()

	var elections *Elections
	err = json.NewDecoder(resp.Body).Decode(&elections)
	if err != nil {
		return &Elections{}, err
	}
	return elections, nil
}

// Challenges Fetches challenges that hotspots owned by the given account are involved in as a challenger, challengee, or witness.
func (a *Account) Challenges(accountID string) (*Challenges, error) {
	resp, err := a.c.Request(http.MethodGet, fmt.Sprintf("/accounts/%s/challenges", accountID), new(bytes.Buffer), nil)
	if err != nil {
		return &Challenges{}, err
	}
	defer resp.Body.Close()

	var challenges *Challenges
	err = json.NewDecoder(resp.Body).Decode(&challenges)
	if err != nil {
		return &Challenges{}, err
	}
	return challenges, nil
}

// PendingTransactions Fetches the outstanding transactions for the given account.
func (a *Account) PendingTransactions(accountID string) (*PendingTransactions, error) {
	resp, err := a.c.Request(http.MethodGet, fmt.Sprintf("/accounts/%s/pending_transactions", accountID), new(bytes.Buffer), nil)
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

// Rewards Returns reward entries by block and gateway for a given account in a timeframe.
func (a *Account) Rewards(accountID string) (*Rewards, error) {
	resp, err := a.c.Request(http.MethodGet, fmt.Sprintf("/accounts/%s/rewards", accountID), new(bytes.Buffer), nil)
	if err != nil {
		return &Rewards{}, err
	}
	defer resp.Body.Close()

	var rewards *Rewards
	err = json.NewDecoder(resp.Body).Decode(&rewards)
	if err != nil {
		return &Rewards{}, err
	}
	return rewards, nil
}

//RewardSum Returns the total rewards for a given account in a given timeframe.
func (a *Account) RewardSum(accountID string) (*RewardSum, error) {
	resp, err := a.c.Request(http.MethodGet, fmt.Sprintf("/accounts/%s/rewards/sum", accountID), new(bytes.Buffer), nil)
	if err != nil {
		return &RewardSum{}, err
	}
	defer resp.Body.Close()

	var rewardSum *RewardSum
	err = json.NewDecoder(resp.Body).Decode(&rewardSum)
	if err != nil {
		return &RewardSum{}, err
	}
	return rewardSum, nil
}

// Stats Fetches account statistics for a given account.
func (a *Account) Stats(accountID string) (*AccountStats, error) {
	resp, err := a.c.Request(http.MethodGet, fmt.Sprintf("/accounts/%s/stats", accountID), new(bytes.Buffer), nil)
	if err != nil {
		return &AccountStats{}, err
	}
	defer resp.Body.Close()
	
	var stats *AccountStats
	err = json.NewDecoder(resp.Body).Decode(&stats)
	if err != nil {
		return &AccountStats{}, err
	}
	return stats, nil
}

func (a *Accounts) Next() bool {
	return len(a.Cursor) > 0
}
