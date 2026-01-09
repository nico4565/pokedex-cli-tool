package main

import (
	"fmt"
)

func commandMap(c *config) error {

	locationAResponse, err := c.httpClient.GetLocationAreaList(c.nextUrl)
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
