package main

import (
	"errors"
	"fmt"
)

func commandMapBack(c *config) error {

	if c.previousUrl == nil {
		return errors.New("you're on the first page")
	}

	locationAResponse, err := c.httpClient.GetLocationAreaList(c.previousUrl)
	if err != nil {
		return err
	}

	for _, locationArea := range locationAResponse.Results {
		fmt.Println(locationArea.Name)
	}

	c.nextUrl = locationAResponse.Next
	c.previousUrl = locationAResponse.Previous

	return err
}
