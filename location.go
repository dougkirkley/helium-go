package helium

import (
	"bytes"
	"fmt"
	"net/http"
	"encoding/json"
)

// Location handles api endpoint /locations docs located at https://docs.helium.com/api/blockchain/locations
type Location struct {
	c *Client
}

// Location returns the Location client
func (c *Client) Location() *Location {
	return &Location{c}
}

type LocationInfo struct {
	Data LocationData `json:"data"`
}

type LocationData struct {
	CityID       string `json:"city_id"`
	Location     string `json:"location"`
	LongCity     string `json:"long_city"`
	LongCountry  string `json:"long_country"`
	LongState    string `json:"long_state"`
	LongStreet   string `json:"long_street"`
	ShortCity    string `json:"short_city"`
	ShortCountry string `json:"short_country"`
	ShortState   string `json:"short_state"`
	ShortStreet  string `json:"short_street"`
}

// Get gets geographic information for a given location
func (l *Location) Get(id string) (*LocationInfo, error) {
	resp, err := l.c.Request(http.MethodGet, fmt.Sprintf("/location/%s", id), new(bytes.Buffer), nil)
	if err != nil {
		return &LocationInfo{}, err
	}
	//defer resp.Body.Close()

	var locationInfo *LocationInfo
	err = json.NewDecoder(resp.Body).Decode(&locationInfo)
	if err != nil {
		return &LocationInfo{}, err
	}
	return locationInfo, nil
}