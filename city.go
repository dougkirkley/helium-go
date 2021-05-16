package helium

import (
	"bytes"
	"fmt"
	"net/http"
	"encoding/json"
)

// City handles api endpoint /cities docs located at https://docs.helium.com/api/blockchain/cities
type City struct {
	c *Client
}

// City returns the City client
func (c *Client) City() *City {
	return &City{c}
}

type Cities struct {
	Data []CityData`json:"data"`
}

type CityData struct {
	CityID       string `json:"city_id"`
	HotspotCount int    `json:"hotspot_count"`
	LongCity     string `json:"long_city"`
	LongCountry  string `json:"long_country"`
	LongState    string `json:"long_state"`
	ShortCity    string `json:"short_city"`
	ShortCountry string `json:"short_country"`
	ShortState   string `json:"short_state"`
}

// Search List all known hotspot cities with the total hotspot count for each city. 
func (c *City) Search(term string) (*Cities, error) {
	if len(term) < 1 {
		return &Cities{}, fmt.Errorf("search term must be 1 character or more, 3 is recommended")
	}
	params := make(map[string]string)
	params["search"] = term
	resp, err := c.c.Request(http.MethodGet, "/cities", new(bytes.Buffer), params)
	if err != nil {
		return &Cities{}, err
	}
	defer resp.Body.Close()

	var cities *Cities
	err = json.NewDecoder(resp.Body).Decode(&cities)
	if err != nil {
		return &Cities{}, err
	}
	return cities, nil
}

// Hotspots Lists all known hotspots for a given city_id.
func (c *City) Hotspots(id string) (*Hotspots, error) {
	resp, err := c.c.Request(http.MethodGet, fmt.Sprintf("/cities/%s/hotspots", id), new(bytes.Buffer), nil)
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