package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(c *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("Command catch requires just one pokemon (id or name) argument! More than one (or maybe zero) arguments found, avoid unrequired spaces in your location area es. area 1 instead of area-1")
	}

	pokemonName := args[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	PokemonResponse, err := c.httpClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	baseExp := PokemonResponse.BaseExperience

	catchResult := rand.Intn(baseExp)

	if catchResult > (baseExp / 2) {
		fmt.Printf("%s was caught!\n", pokemonName)
		if _, exists := c.caughtPokemon[pokemonName]; !exists {
			c.caughtPokemon[pokemonName] = PokemonResponse
		}
		return err
	}

	fmt.Printf("%s escaped!\n", pokemonName)
	return err
}
