package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (ResponseShallowLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ResponseShallowLocations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return ResponseShallowLocations{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return ResponseShallowLocations{}, err
	}

	locationsResponse := ResponseShallowLocations{}
	err = json.Unmarshal(dat, &locationsResponse)

	if err != nil {
		return ResponseShallowLocations{}, err
	}

	return locationsResponse, nil
}
