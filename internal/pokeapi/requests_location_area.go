package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func (c *Client) GetLocationAreaList(urlPtr *string) (LocationArea, error) {

	url := baseURL + "/location-area"
	if urlPtr != nil {
		url = *urlPtr
	}

	var body []byte

	if val, ok := c.cache.Get(url); ok {
		locationsResp := LocationArea{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return LocationArea{}, err
		}

		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationArea{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}
	defer resp.Body.Close()
	body, err = io.ReadAll(resp.Body)

	if resp.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", resp.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
		return LocationArea{}, err
	}
	c.cache.Add(url, body)

	lAResponse := LocationArea{}
	err = json.Unmarshal(body, &lAResponse)
	if err != nil {
		fmt.Println(err)
	}

	return lAResponse, nil
}

func (c *Client) GetLocationAreaDetailedResponse(nameOrIdLocation string) (LocationAreaDetailed, error) {

	url := baseURL + "/location-area" + "/" + nameOrIdLocation

	var body []byte

	if val, ok := c.cache.Get(url); ok {
		locationsResp := LocationAreaDetailed{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return LocationAreaDetailed{}, err
		}

		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreaDetailed{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaDetailed{}, err
	}
	defer resp.Body.Close()
	body, err = io.ReadAll(resp.Body)

	if resp.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", resp.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
		return LocationAreaDetailed{}, err
	}
	c.cache.Add(url, body)

	response := LocationAreaDetailed{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println(err)
	}

	return response, nil
}
