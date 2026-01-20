package main

import (
	"errors"
	"fmt"
)

func commandExplore(c *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("Command explore requires just one location area (id or name) argument! More than one (or maybe zero) arguments found, avoid unrequired spaces in your location area es. area 1 instead of area-1")
	}
	locationADResponse, err := c.httpClient.GetLocationAreaDetailedResponse(args[0])
	if err != nil {
		return err
	}

	for _, pokemonEncounter := range locationADResponse.PokemonEncounters {
		fmt.Println(pokemonEncounter.Pokemon.Name)
	}

	return err
}
