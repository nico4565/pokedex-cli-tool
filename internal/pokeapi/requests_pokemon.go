package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func (c *Client) GetPokemon(nameOrIdLocation string) (Pokemon, error) {

	url := baseURL + "/pokemon" + "/" + nameOrIdLocation

	var body []byte

	if val, ok := c.cache.Get(url); ok {
		locationsResp := Pokemon{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return Pokemon{}, err
		}

		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer resp.Body.Close()
	body, err = io.ReadAll(resp.Body)

	if resp.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", resp.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
		return Pokemon{}, err
	}
	c.cache.Add(url, body)

	response := Pokemon{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println(err)
	}

	return response, nil
}
