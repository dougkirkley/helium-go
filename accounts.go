package helium

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Account struct {
	client *Client
}

type Accounts struct {
	Data   []AccountData `json:"data"`
	Cursor string        `json:"cursor"`
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

type Hotspots struct {
	Data []HotspotData `json:"data"`
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
type HotspotData struct {
	Lng              float64   `json:"lng"`
	Lat              float64   `json:"lat"`
	TimestampAdded   time.Time `json:"timestamp_added"`
	Status           Status    `json:"status"`
	RewardScale      float64   `json:"reward_scale"`
	Owner            string    `json:"owner"`
	Nonce            int       `json:"nonce"`
	Name             string    `json:"name"`
	Location         string    `json:"location"`
	LastPocChallenge int       `json:"last_poc_challenge"`
	LastChangeBlock  int       `json:"last_change_block"`
	Geocode          Geocode   `json:"geocode"`
	Gain             int       `json:"gain"`
	Elevation        int       `json:"elevation"`
	BlockAdded       int       `json:"block_added"`
	Block            int       `json:"block"`
	Address          string    `json:"address"`
}

type Ouis struct {
	Data []OuiData `json:"data"`
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
	Data []ActivityData `json:"data"`
}
type Rewards struct {
	Type    string `json:"type"`
	Gateway string `json:"gateway"`
	Amount  int    `json:"amount"`
	Account string `json:"account"`
}
type ActivityData struct {
	Type       string    `json:"type"`
	Time       int       `json:"time"`
	StartEpoch int       `json:"start_epoch"`
	Rewards    []Rewards `json:"rewards"`
	Height     int       `json:"height"`
	Hash       string    `json:"hash"`
	EndEpoch   int       `json:"end_epoch"`
}

type ActivityCount struct {
	Data CountsData `json:"data"`
}
type CountsData struct {
	AddGatewayV1     int `json:"add_gateway_v1"`
	AssertLocationV1 int `json:"assert_location_v1"`
}

func (c *Client) Account() *Account {
	return &Account{c}
}

func (a *Account) List(cursor string) (*Accounts, error) {
	params := make(map[string]string)
	if len(cursor) > 0 {
		params["cursor"] = cursor
	}
	resp, err := a.client.Request(http.MethodGet, "/accounts", params)
	if err != nil {
		return &Accounts{}, err
	}
	var accounts *Accounts
	err = json.Unmarshal(resp, &accounts)
	if err != nil {
		return &Accounts{}, err
	}
	return accounts, nil
}

func (a *Account) Richest(limit int) (*Accounts, error) {
	params := make(map[string]string)
	if limit > 0 {
		params["limit"] = fmt.Sprintf("%v", limit)
	}
	resp, err := a.client.Request(http.MethodGet, "/accounts/rich", params)
	if err != nil {
		return &Accounts{}, err
	}
	var accounts *Accounts
	err = json.Unmarshal(resp, &accounts)
	if err != nil {
		return &Accounts{}, err
	}
	return accounts, nil
}

func (a *Account) Get(accountID string) (*AccountData, error) {
	resp, err := a.client.Request(http.MethodGet, fmt.Sprintf("/accounts/%s", accountID), nil)
	if err != nil {
		return &AccountData{}, err
	}
	var account *AccountData
	err = json.Unmarshal(resp, &account)
	if err != nil {
		return &AccountData{}, err
	}
	return account, nil
}

func (a *Account) Hotspots(accountID string) (*Hotspots, error) {
	resp, err := a.client.Request(http.MethodGet, fmt.Sprintf("/accounts/%s/hotspots", accountID), nil)
	if err != nil {
		return &Hotspots{}, err
	}
	var hotspots *Hotspots
	err = json.Unmarshal(resp, &hotspots)
	if err != nil {
		return &Hotspots{}, err
	}
	return hotspots, nil
}

func (a *Account) Ouis(accountID string) (*Ouis, error) {
	resp, err := a.client.Request(http.MethodGet, fmt.Sprintf("/accounts/%s/ouis", accountID), nil)
	if err != nil {
		return &Ouis{}, err
	}
	var ouis *Ouis
	err = json.Unmarshal(resp, &ouis)
	if err != nil {
		return &Ouis{}, err
	}
	return ouis, nil
}

func (a *Account) Activity(accountID string) (*Activity, error) {
	resp, err := a.client.Request(http.MethodGet, fmt.Sprintf("/accounts/%s/activity", accountID), nil)
	if err != nil {
		return &Activity{}, err
	}
	var activity *Activity
	err = json.Unmarshal(resp, &activity)
	if err != nil {
		return &Activity{}, err
	}
	return activity, nil
}

func (a *Account) ActivityCount(accountID string) (*ActivityCount, error) {
	resp, err := a.client.Request(http.MethodGet, fmt.Sprintf("/accounts/%s/activity/count", accountID), nil)
	if err != nil {
		return &ActivityCount{}, err
	}
	var activityCount *ActivityCount
	err = json.Unmarshal(resp, &activityCount)
	if err != nil {
		return &ActivityCount{}, err
	}
	return activityCount, nil
}

func (a *Accounts) Next() bool {
	return len(a.Cursor) > 0
}
