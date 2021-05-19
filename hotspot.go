package helium

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// Hotspot handles api endpoint /hotspots docs located at https://docs.helium.com/api/blockchain/hotspots
type Hotspot struct {
	c *Client
}

// Hotspot returns the Hotspot client
func (c *Client) Hotspot() *Hotspot {
	return &Hotspot{c}
}

type Hotspots struct {
	Data   []HotspotData `json:"data"`
	Cursor string        `json:"cursor"`
}

type HotspotsData struct {
	Lng               float64 `json:"lng"`
	Lat               float64 `json:"lat"`
	Status            Status  `json:"status"`
	ScoreUpdateHeight int     `json:"score_update_height"`
	Score             float64 `json:"score"`
	Owner             string  `json:"owner"`
	Nonce             int     `json:"nonce"`
	Name              string  `json:"name"`
	Location          string  `json:"location"`
	Geocode           Geocode `json:"geocode"`
	BlockAdded        int     `json:"block_added"`
	Block             int     `json:"block"`
	Address           string  `json:"address"`
}

type HotspotInfo struct {
	Data HotspotData `json:"data"`
}

type HotspotData struct {
	Address           string  `json:"address"`
	Block             int     `json:"block"`
	BlockAdded        int     `json:"block_added"`
	Geocode           Geocode `json:"geocode"`
	Lat               float64 `json:"lat"`
	Lng               float64 `json:"lng"`
	Location          string  `json:"location"`
	Name              string  `json:"name"`
	Nonce             int     `json:"nonce"`
	Owner             string  `json:"owner"`
	Score             float64 `json:"score"`
	ScoreUpdateHeight int     `json:"score_update_height"`
	Status            Status  `json:"status"`
}

type HotspotsActivity struct {
	Data []HotspotsActivityData `json:"data"`
}

type HotspotsActivityData struct {
	Fee        int     `json:"fee"`
	Gateway    string  `json:"gateway"`
	Hash       string  `json:"hash"`
	Height     int     `json:"height"`
	Lat        float64 `json:"lat,omitempty"`
	Lng        float64 `json:"lng,omitempty"`
	Location   string  `json:"location,omitempty"`
	Nonce      int     `json:"nonce,omitempty"`
	Owner      string  `json:"owner"`
	Payer      string  `json:"payer"`
	StakingFee int     `json:"staking_fee"`
	Time       int     `json:"time"`
	Type       string  `json:"type"`
}

type HotspotActivityCount struct {
	Data HotspotActivityCountData `json:"data"`
}

type HotspotActivityCountData struct {
	AssertLocationV1 int `json:"assert_location_v1"`
}

type Witnesses struct {
	Data []WitnessData `json:"data"`
}

type WitnessData struct {
	Address           string      `json:"address"`
	Block             int         `json:"block"`
	BlockAdded        int         `json:"block_added"`
	Geocode           Geocode     `json:"geocode"`
	Lat               float64     `json:"lat"`
	Lng               float64     `json:"lng"`
	Location          string      `json:"location"`
	Name              string      `json:"name"`
	Nonce             int         `json:"nonce"`
	Owner             string      `json:"owner"`
	Score             float64     `json:"score"`
	ScoreUpdateHeight int         `json:"score_update_height"`
	Status            Status      `json:"status"`
	WitnessFor        string      `json:"witness_for"`
	WitnessInfo       WitnessInfo `json:"witness_info"`
}

type Histogram struct {
}

type WitnessInfo struct {
	FirstTime  int64     `json:"first_time"`
	Histogram  Histogram `json:"histogram"`
	RecentTime int64     `json:"recent_time"`
}

type HotspotInput struct {
	Address string
	Name    string
}

type HotspotSearchInput struct {
	Term string
}

type HotspotDistanceInput struct {
	Lat      float64
	Lon      float64
	Distance int
}

type HotspotBoxInput struct {
	Swlat float64
	Swlon float64
	Nelat float64
	Nelon float64
}

type HotspotHexInput struct {
	ID string
}

type HotspotRewardsInput struct {
	Address string
	MaxTime string
	MinTime string
}

// List known hotspots as registered on the blockchain.
func (h *Hotspot) List() (*Hotspots, error) {
	resp, err := h.c.Request(http.MethodGet, "/hotspots", new(bytes.Buffer), nil)
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

// Get Fetch a hotspot with a given address.
func (h *Hotspot) Get(input *HotspotInput) (*HotspotInfo, error) {
	resp, err := h.c.Request(http.MethodGet, fmt.Sprintf("/hotspots/%s", input.Address), new(bytes.Buffer), nil)
	if err != nil {
		return &HotspotInfo{}, err
	}
	defer resp.Body.Close()

	var hotspotInfo *HotspotInfo
	err = json.NewDecoder(resp.Body).Decode(&hotspotInfo)
	if err != nil {
		return &HotspotInfo{}, err
	}
	return hotspotInfo, nil
}

// GetByName Fetch the hotspots which map to the given 3-word animal name.
func (h *Hotspot) GetByName(input *HotspotInput) (*Hotspots, error) {
	resp, err := h.c.Request(http.MethodGet, fmt.Sprintf("/hotspots/name/%s", input.Name), new(bytes.Buffer), nil)
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

// Search Fetch the hotspots which match a search term in the given search term query parameter.
func (h *Hotspot) Search(input *HotspotSearchInput) (*Hotspots, error) {
	if len(input.Term) < 1 {
		return &Hotspots{}, fmt.Errorf("search term must be 1 character or more, 3 is recommended")
	}
	params := make(map[string]string)
	params["search"] = input.Term
	resp, err := h.c.Request(http.MethodGet, "/hotspots/name", new(bytes.Buffer), params)
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

// Distance Fetch the hotspots which are within a given number of meters from the given lat and lon coordinates.
func (h *Hotspot) Distance(input *HotspotDistanceInput) (*Hotspots, error) {
	params := make(map[string]string)
	params["lat"] = fmt.Sprintf("%v", input.Lat)
	params["lon"] = fmt.Sprintf("%v", input.Lon)
	params["distance"] = fmt.Sprintf("%v", input.Distance)
	resp, err := h.c.Request(http.MethodGet, "/hotspots/location/distance", new(bytes.Buffer), params)
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

// Box Fetch the hotspots which are within a given geographic boundary indicated by it's south-wesetern and north-eastern co-ordinates.
func (h *Hotspot) Box(input *HotspotBoxInput) (*Hotspots, error) {
	params := make(map[string]string)
	params["swlat"] = fmt.Sprintf("%v", input.Swlat)
	params["swlon"] = fmt.Sprintf("%v", input.Swlon)
	params["nelat"] = fmt.Sprintf("%v", input.Nelat)
	params["nelon"] = fmt.Sprintf("%v", input.Nelon)
	resp, err := h.c.Request(http.MethodGet, "/hotspots/location/box", new(bytes.Buffer), params)
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

// GetByHex Fetch the hotspots which are in the given h3 index.
func (h *Hotspot) GetByHex(input *HotspotHexInput) (*HotspotInfo, error) {
	resp, err := h.c.Request(http.MethodGet, fmt.Sprintf("/hotspots/hex/%s", input.ID), new(bytes.Buffer), nil)
	if err != nil {
		return &HotspotInfo{}, err
	}
	defer resp.Body.Close()
	var hotspotInfo *HotspotInfo
	err = json.NewDecoder(resp.Body).Decode(&hotspotInfo)
	if err != nil {
		return &HotspotInfo{}, err
	}
	return hotspotInfo, nil
}

// Activity Lists all blockchain transactions that the given hotspot was involved in.
func (h *Hotspot) Activity(input *HotspotInput) (*HotspotsActivity, error) {
	resp, err := h.c.Request(http.MethodGet, fmt.Sprintf("/hotspots/%s/activity", input.Address), new(bytes.Buffer), nil)
	if err != nil {
		return &HotspotsActivity{}, err
	}
	defer resp.Body.Close()

	var hotspotsActivity *HotspotsActivity
	err = json.NewDecoder(resp.Body).Decode(&hotspotsActivity)
	if err != nil {
		return &HotspotsActivity{}, err
	}
	return hotspotsActivity, nil
}

// ActivityCount Count transactions that indicate activity for a hotspot.
func (h *Hotspot) ActivityCount(input *HotspotInput) (*HotspotActivityCount, error) {
	resp, err := h.c.Request(http.MethodGet, fmt.Sprintf("/hotspots/%s/activity/count", input.Address), new(bytes.Buffer), nil)
	if err != nil {
		return &HotspotActivityCount{}, err
	}
	defer resp.Body.Close()

	var hotspotActivityCount *HotspotActivityCount
	err = json.NewDecoder(resp.Body).Decode(&hotspotActivityCount)
	if err != nil {
		return &HotspotActivityCount{}, err
	}
	return hotspotActivityCount, nil
}

// Elections Lists the consensus group transactions that the given hotspot was involved in.
func (h *Hotspot) Elections(input *HotspotInput) (*Elections, error) {
	resp, err := h.c.Request(http.MethodGet, fmt.Sprintf("/hotspots/%s/elections", input.Address), new(bytes.Buffer), nil)
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

// CurrentlyElected Returns the list of hotspots that are currently elected to the consensus group.
func (h *Hotspot) CurrentlyElected() (*Elections, error) {
	resp, err := h.c.Request(http.MethodGet, "/hotspots/elected", new(bytes.Buffer), nil)
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

// Challenges Lists the challenge (receipts) that the given hotspot a challenger, challengee or a witness in.
func (h *Hotspot) Challenges(input *HotspotInput) (*Challenges, error) {
	resp, err := h.c.Request(http.MethodGet, fmt.Sprintf("/hotspots/%s/challenges", input.Address), new(bytes.Buffer), nil)
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

// Rewards Returns reward entries by block and gateway for a given account in a timeframe.
func (h *Hotspot) Rewards(input *HotspotRewardsInput) (*Rewards, error) {
	params := make(map[string]string)
	params["min_time"] = input.MinTime
	params["max_time"] = input.MaxTime
	resp, err := h.c.Request(http.MethodGet, fmt.Sprintf("/hotspots/%s/rewards", input.Address), new(bytes.Buffer), params)
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

//RewardSum Returns rewards for a given hotspot per reward block the hotspot is in, for a given timeframe.
func (h *Hotspot) RewardSum(input *HotspotInput) (*RewardSum, error) {
	resp, err := h.c.Request(http.MethodGet, fmt.Sprintf("/hotspots/%s/rewards/sum", input.Address), new(bytes.Buffer), nil)
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
