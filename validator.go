package helium

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Validator handles api endpoint /validators docs located at https://docs.helium.com/api/blockchain/validators
type Validator struct {
	c *Client
}

// Validator returns the Validator client
func (c *Client) Validator() *Validator {
	return &Validator{c}
}

// Validators holds the data for the /validators endpoint
type Validators struct {
	Data   []ValidatorData `json:"data"`
	Cursor string          `json:"cursor"`
}

// ValidatorData the data per validator
type ValidatorData struct {
	VersionHeartbeat int           `json:"version_heartbeat"`
	Status           Status        `json:"status"`
	StakeStatus      string        `json:"stake_status"`
	Stake            int64         `json:"stake"`
	Penalty          int           `json:"penalty"`
	Penalties        []interface{} `json:"penalties"`
	Owner            string        `json:"owner"`
	Name             string        `json:"name"`
	LastHeartbeat    int           `json:"last_heartbeat"`
	BlockAdded       int           `json:"block_added"`
	Block            int           `json:"block"`
	Address          string        `json:"address"`
}

// ValidatorInfo info for a singular validator
type ValidatorInfo struct {
	Data ValidatorData `json:"data"`
}

// Penalties penalty info
type Penalties struct {
	Type   string `json:"type"`
	Height int    `json:"height"`
	Amount int    `json:"amount"`
}

// ValidatorActivity holds the data for /validators/:validator/activity endpoint
type ValidatorActivity struct {
	Cursor string                  `json:"cursor"`
	Data   []ValidatorActivityData `json:"data"`
}

// ValidatorActivityData the data per validator activity
type ValidatorActivityData struct {
	Address   string `json:"address"`
	Hash      string `json:"hash"`
	Height    int    `json:"height"`
	Signature string `json:"signature"`
	Time      int    `json:"time"`
	Type      string `json:"type"`
	Version   int    `json:"version"`
}

// ValidatorActivityCount holds the data for /validators/:validator/activity/count endpoint
type ValidatorActivityCount struct {
	Data ValidatorActivityCountData `json:"data"`
}

// ValidatorActivityCountsData the data per validator activity count
type ValidatorActivityCountData struct {
	VarsV1                   int `json:"vars_v1"`
	ValidatorHeartbeatV1     int `json:"validator_heartbeat_v1"`
	UnstakeValidatorV1       int `json:"unstake_validator_v1"`
	TransferValidatorStakeV1 int `json:"transfer_validator_stake_v1"`
	TransferHotspotV1        int `json:"transfer_hotspot_v1"`
	TokenBurnV1              int `json:"token_burn_v1"`
	TokenBurnExchangeRateV1  int `json:"token_burn_exchange_rate_v1"`
	StateChannelOpenV1       int `json:"state_channel_open_v1"`
	StateChannelCloseV1      int `json:"state_channel_close_v1"`
	StakeValidatorV1         int `json:"stake_validator_v1"`
	SecurityExchangeV1       int `json:"security_exchange_v1"`
	SecurityCoinbaseV1       int `json:"security_coinbase_v1"`
	RoutingV1                int `json:"routing_v1"`
	RewardsV2                int `json:"rewards_v2"`
	RewardsV1                int `json:"rewards_v1"`
	RedeemHtlcV1             int `json:"redeem_htlc_v1"`
	PriceOracleV1            int `json:"price_oracle_v1"`
	PocRequestV1             int `json:"poc_request_v1"`
	PocReceiptsV1            int `json:"poc_receipts_v1"`
	PaymentV2                int `json:"payment_v2"`
	PaymentV1                int `json:"payment_v1"`
	OuiV1                    int `json:"oui_v1"`
	GenGatewayV1             int `json:"gen_gateway_v1"`
	DcCoinbaseV1             int `json:"dc_coinbase_v1"`
	CreateHtlcV1             int `json:"create_htlc_v1"`
	ConsensusGroupV1         int `json:"consensus_group_v1"`
	ConsensusGroupFailureV1  int `json:"consensus_group_failure_v1"`
	CoinbaseV1               int `json:"coinbase_v1"`
	AssertLocationV2         int `json:"assert_location_v2"`
	AssertLocationV1         int `json:"assert_location_v1"`
	AddGatewayV1             int `json:"add_gateway_v1"`
}

// ValidatorStats holds the data for /validators/:validator/stats
type ValidatorStats struct {
	Data ValidatorStatsData `json:"data"`
}

// Cooldown struct
type Cooldown struct {
	Amount int `json:"amount"`
	Count  int `json:"count"`
}

// Stakes struct
type Staked struct {
	Amount float64 `json:"amount"`
	Count  int     `json:"count"`
}

// Unstaked
type Unstaked struct {
	Amount float64 `json:"amount"`
	Count  int     `json:"count"`
}

// ValidatorStatsData data for validator stats
type ValidatorStatsData struct {
	Active   int      `json:"active"`
	Cooldown Cooldown `json:"cooldown"`
	Staked   Staked   `json:"staked"`
	Unstaked Unstaked `json:"unstaked"`
}

// ValidatorElections holds the data for the /validators/:validator/elected endpoint
type ValidatorElections struct {
	Data []ValidatorElectionData `json:"data"`
}

// ValidatorelectionData data for validator election
type ValidatorElectionData struct {
	VersionHeartbeat int    `json:"version_heartbeat"`
	Status           string `json:"status"`
	Stake            int64  `json:"stake"`
	Owner            string `json:"owner"`
	LastHeartbeat    int    `json:"last_heartbeat"`
	Block            int    `json:"block"`
	Address          string `json:"address"`
}

// ValidatorRewards holds the data for the /validator/:validator/rewards endpoint
type ValidatorRewards struct {
	Data []ValidatorRewardData `json:"data"`
}

// ValidatorRewardData data for validator reward
type ValidatorRewardData struct {
	Account   string    `json:"account"`
	Amount    int       `json:"amount"`
	Block     int       `json:"block"`
	Gateway   string    `json:"gateway"`
	Hash      string    `json:"hash"`
	Timestamp time.Time `json:"timestamp"`
}

// ValidatorRewardsSum holds the data for the /validator/:validator/rewards/sum endpoint
type ValidatorRewardsSum struct {
	Data ValidatorRewardsSumData `json:"data"`
	Meta Meta `json:"meta"`
}

// ValidatorRewardsSumData data for validator rewards sum
type ValidatorRewardsSumData struct {
	Avg    float64 `json:"avg"`
	Max    float64 `json:"max"`
	Median float64 `json:"median"`
	Min    float64 `json:"min"`
	Stddev float64 `json:"stddev"`
	Sum    int64   `json:"sum"`
	Total  float64 `json:"total"`
}

// Meta struct
type Meta struct {
	MaxTime time.Time `json:"max_time"`
	MinTime time.Time `json:"min_time"`
}

// List List known validators as registered on the blockchain.
func (v *Validator) List() (*Validators, error) {
	resp, err := v.c.Request(http.MethodGet, "/validators", new(bytes.Buffer), nil)
	if err != nil {
		return &Validators{}, err
	}
	defer resp.Body.Close()

	var validators *Validators
	err = json.NewDecoder(resp.Body).Decode(&validators)
	if err != nil {
		return &Validators{}, err
	}
	return validators, nil
}

// Get Fetch a validator with a given address.
func (v *Validator) Get(address string) (*ValidatorInfo, error) {
	resp, err := v.c.Request(http.MethodGet, fmt.Sprintf("/validators/%s", address), new(bytes.Buffer), nil)
	if err != nil {
		return &ValidatorInfo{}, err
	}
	defer resp.Body.Close()

	var validatorInfo *ValidatorInfo
	err = json.NewDecoder(resp.Body).Decode(&validatorInfo)
	if err != nil {
		return &ValidatorInfo{}, err
	}
	return validatorInfo, nil
}

// GetByName Fetch the validators which map to the given 3-word animal name. 
func (v *Validator) GetByName(name string) (*ValidatorInfo, error) {
	resp, err := v.c.Request(http.MethodGet, fmt.Sprintf("/validators/name/%s", name), new(bytes.Buffer), nil)
	if err != nil {
		return &ValidatorInfo{}, err
	}
	defer resp.Body.Close()

	var validatorInfo *ValidatorInfo
	err = json.NewDecoder(resp.Body).Decode(&validatorInfo)
	if err != nil {
		return &ValidatorInfo{}, err
	}
	return validatorInfo, nil
}

// Search Fetch the validators which match a search term in the given search term query parameter.
func (v *Validator) Search(term string) (*Validators, error) {
	if len(term) < 1 {
		return &Validators{}, fmt.Errorf("search term must be 1 character or more, 3 is recommended")
	}
	params := make(map[string]string)
	params["search"] = term
	resp, err := v.c.Request(http.MethodGet, "/validators/name", new(bytes.Buffer), params)
	if err != nil {
		return &Validators{}, err
	}
	defer resp.Body.Close()

	var validators *Validators
	err = json.NewDecoder(resp.Body).Decode(&validators)
	if err != nil {
		return &Validators{}, err
	}
	return validators, nil
}

// Activity Lists all blockchain transactions that the given validator was involved in.
func (v *Validator) Activity(address string) (*ValidatorActivity, error) {
	resp, err := v.c.Request(http.MethodGet, fmt.Sprintf("/validators/%s/activity", address), new(bytes.Buffer), nil)
	if err != nil {
		return &ValidatorActivity{}, err
	}
	defer resp.Body.Close()

	var validatorActivity *ValidatorActivity
	err = json.NewDecoder(resp.Body).Decode(&validatorActivity)
	if err != nil {
		return &ValidatorActivity{}, err
	}
	return validatorActivity, nil
}

// ActivityCount Count transactions that indicate activity for a validator.
func (v *Validator) ActivityCount(address string, filterTypes string) (*ValidatorActivityCount, error) {
	params := make(map[string]string)
	if len(filterTypes) > 0 {
		params["filter_types"] = filterTypes
	}
	resp, err := v.c.Request(http.MethodGet, fmt.Sprintf("/validators/%s/activity/count", address), new(bytes.Buffer), params)
	if err != nil {
		return &ValidatorActivityCount{}, err
	}
	defer resp.Body.Close()

	var validatorActivityCount *ValidatorActivityCount
	err = json.NewDecoder(resp.Body).Decode(&validatorActivityCount)
	if err != nil {
		return &ValidatorActivityCount{}, err
	}
	return validatorActivityCount, nil
}

// Stats Returns stats for validators
func (v *Validator) Stats(address string) (*ValidatorStats, error) {
	resp, err := v.c.Request(http.MethodGet, "/validators/stats", new(bytes.Buffer), nil)
	if err != nil {
		return &ValidatorStats{}, err
	}
	defer resp.Body.Close()

	var validatorStats *ValidatorStats
	err = json.NewDecoder(resp.Body).Decode(&validatorStats)
	if err != nil {
		return &ValidatorStats{}, err
	}
	return validatorStats, nil
}

// ListElected Returns the list of validators that are currently elected to the consensus group.
func (v *Validator) ListElected() (*ValidatorElections, error) {
	resp, err := v.c.Request(http.MethodGet, "/validators/elected", new(bytes.Buffer), nil)
	if err != nil {
		return &ValidatorElections{}, err
	}
	defer resp.Body.Close()

	var validatorElections *ValidatorElections
	err = json.NewDecoder(resp.Body).Decode(&validatorElections)
	if err != nil {
		return &ValidatorElections{}, err
	}
	return validatorElections, nil
}

// ElectedAtHeight Returns the list of validators that were in the consensus group at a given block height
func (v *Validator) ElectedAtHeight(height string,) (*Validators, error) {
	resp, err := v.c.Request(http.MethodGet, fmt.Sprintf("/validators/elected/%s", height), new(bytes.Buffer), nil)
	if err != nil {
		return &Validators{}, err
	}
	defer resp.Body.Close()

	var validators *Validators
	err = json.NewDecoder(resp.Body).Decode(&validators)
	if err != nil {
		return &Validators{}, err
	}
	return validators, nil
}

// ElectedAtHash Returns the list of validators that were elected in the consensus group transcation indicated by the given transaction hash.
func (v *Validator) ElectedAtHash(hash string,) (*Validators, error) {
	resp, err := v.c.Request(http.MethodGet, fmt.Sprintf("/validators/elected/hash/%s", hash), new(bytes.Buffer), nil)
	if err != nil {
		return &Validators{}, err
	}
	defer resp.Body.Close()

	var validators *Validators
	err = json.NewDecoder(resp.Body).Decode(&validators)
	if err != nil {
		return &Validators{}, err
	}
	return validators, nil
}

// Rewards Returns rewards for a given validator per reward block the validator is in, for a given timeframe. 
func (v *Validator) Rewards(address string, cursor string, maxTime string, minTime string) (*Validators, error) {
	params := make(map[string]string)
	if len(cursor) > 0 {
		params["cursor"] = cursor
	}
	params["max_time"] = maxTime
	params["min_time"] = minTime

	resp, err := v.c.Request(http.MethodGet, fmt.Sprintf("/validators/%s/rewards", address), new(bytes.Buffer), params)
	if err != nil {
		return &Validators{}, err
	}
	defer resp.Body.Close()

	var validators *Validators
	err = json.NewDecoder(resp.Body).Decode(&validators)
	if err != nil {
		return &Validators{}, err
	}
	return validators, nil
}

// RewardsSum Returns the total rewards earned for a given validator over a given time range.
func (v *Validator) RewardsSum(address string,) (*ValidatorRewardsSum, error) {
	resp, err := v.c.Request(http.MethodGet, fmt.Sprintf("/validators/%s/rewards/sum", address), new(bytes.Buffer), nil)
	if err != nil {
		return &ValidatorRewardsSum{}, err
	}
	defer resp.Body.Close()
	
	var validatorRewardsSum *ValidatorRewardsSum
	err = json.NewDecoder(resp.Body).Decode(&validatorRewardsSum)
	if err != nil {
		return &ValidatorRewardsSum{}, err
	}
	return validatorRewardsSum, nil
}