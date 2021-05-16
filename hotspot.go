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
func (h *Hotspot) Get(address string) (*HotspotInfo, error) {
	resp, err := h.c.Request(http.MethodGet, fmt.Sprintf("/hotspots/%s", address), new(bytes.Buffer), nil)
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
func (h *Hotspot) GetByName(name string) (*Hotspots, error) {
	resp, err := h.c.Request(http.MethodGet, fmt.Sprintf("/hotspots/name/%s", name), new(bytes.Buffer), nil)
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
func (h *Hotspot) Search(term string) (*Hotspots, error) {
	if len(term) < 1 {
		return &Hotspots{}, fmt.Errorf("search term must be 1 character or more, 3 is recommended")
	}
	params := make(map[string]string)
	params["search"] = term
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
func (h *Hotspot) Distance(lat float64, lon float64, distance int) (*Hotspots, error) {
	params := make(map[string]string)
	params["lat"] = fmt.Sprintf("%v", lat)
	params["lon"] = fmt.Sprintf("%v", lon)
	params["distance"] = fmt.Sprintf("%v", distance)
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
func (h *Hotspot) Box(swlat float64, swlon float64, nelat float64, nelon float64) (*Hotspots, error) {
	params := make(map[string]string)
	params["swlat"] = fmt.Sprintf("%v", swlat)
	params["swlon"] = fmt.Sprintf("%v", swlon)
	params["nelat"] = fmt.Sprintf("%v", nelat)
	params["nelon"] = fmt.Sprintf("%v", nelon)
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
func (h *Hotspot) GetByHex(hex string) (*HotspotInfo, error) {
	resp, err := h.c.Request(http.MethodGet, fmt.Sprintf("/hotspots/hex/%s", hex), new(bytes.Buffer), nil)
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
func (h *Hotspot) Activity(address string) (*HotspotsActivity, error) {
	resp, err := h.c.Request(http.MethodGet, fmt.Sprintf("/hotspots/%s/activity", address), new(bytes.Buffer), nil)
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
func (h *Hotspot) ActivityCount(address string) (*HotspotActivityCount, error) {
	resp, err := h.c.Request(http.MethodGet, fmt.Sprintf("/hotspots/%s/activity/count", address), new(bytes.Buffer), nil)
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
func (h *Hotspot) Elections(address string) (*Elections, error) {
	resp, err := h.c.Request(http.MethodGet, fmt.Sprintf("/hotspots/%s/elections", address), new(bytes.Buffer), nil)
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
func (h *Hotspot) Challenges(address string) (*Challenges, error) {
	resp, err := h.c.Request(http.MethodGet, fmt.Sprintf("/hotspots/%s/challenges", address), new(bytes.Buffer), nil)
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
func (h *Hotspot) Rewards(address string, maxTime string, minTime string) (*Rewards, error) {
	params := make(map[string]string)
	params["min_time"] = minTime
	params["max_time"] = maxTime
	resp, err := h.c.Request(http.MethodGet, fmt.Sprintf("/hotspots/%s/rewards", address), new(bytes.Buffer), params)
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
func (h *Hotspot) RewardSum(address string) (*RewardSum, error) {
	resp, err := h.c.Request(http.MethodGet, fmt.Sprintf("/hotspots/%s/rewards/sum", address), new(bytes.Buffer), nil)
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
