package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func (c *Client) GetLocationAreaList(urlPtr *string) (LocationAreaResponse, error) {

	url := baseURL + "/location-area"
	if urlPtr != nil {
		url = *urlPtr
	}

	var body []byte

	if val, ok := c.cache.Get(url); ok {
		locationsResp := LocationAreaResponse{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return LocationAreaResponse{}, err
		}

		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaResponse{}, err
	}
	defer resp.Body.Close()
	body, err = io.ReadAll(resp.Body)

	if resp.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", resp.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
		return LocationAreaResponse{}, err
	}
	c.cache.Add(url, body)

	lAResponse := LocationAreaResponse{}
	err = json.Unmarshal(body, &lAResponse)
	if err != nil {
		fmt.Println(err)
	}

	return lAResponse, nil
}

func (c *Client) GetLocationAreaDetailedResponse(urlPtr *string, nameOrIdLocation string) (LocationAreaDetailedResponse, error) {

	url := baseURL + "/location-area" + "/" + nameOrIdLocation
	if urlPtr != nil {
		url = *urlPtr
	}

	var body []byte

	if val, ok := c.cache.Get(url); ok {
		locationsResp := LocationAreaDetailedResponse{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return LocationAreaDetailedResponse{}, err
		}

		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreaDetailedResponse{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaDetailedResponse{}, err
	}
	defer resp.Body.Close()
	body, err = io.ReadAll(resp.Body)

	if resp.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", resp.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
		return LocationAreaDetailedResponse{}, err
	}
	c.cache.Add(url, body)

	response := LocationAreaDetailedResponse{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println(err)
	}

	return response, nil
}
