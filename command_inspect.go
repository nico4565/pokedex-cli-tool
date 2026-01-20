package main

import (
	"errors"
	"fmt"
)

func commandInspect(c *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("Command inspect requires just one pokemon (id or name) argument! More than one/no argument found, avoid unrequired spaces!")
	}

	pokemonName := args[0]

	pokemon, exists := c.caughtPokemon[pokemonName]
	if !exists {
		fmt.Printf("you have not caught that pokemon\n")
		return nil
	}

	fmt.Println("Name:", pokemon.Name)
	fmt.Println("Height:", pokemon.Height)
	fmt.Println("Weight:", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, typeInfo := range pokemon.Types {
		fmt.Println("  -", typeInfo.Type.Name)
	}
	return nil
}
