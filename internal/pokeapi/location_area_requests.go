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

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaResponse{}, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if resp.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", resp.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
		return LocationAreaResponse{}, err
	}

	lAResponse := LocationAreaResponse{}
	err = json.Unmarshal(body, &lAResponse)
	if err != nil {
		fmt.Println(err)
	}

	return lAResponse, nil
}
