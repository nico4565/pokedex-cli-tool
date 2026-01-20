package main

import (
	"fmt"
)

func commandPokedex(c *config, args ...string) error {
	fmt.Println("Your Pokedex:")
	for k := range c.caughtPokemon {
		fmt.Printf("- %s\n", k)
	}
	return nil
}
