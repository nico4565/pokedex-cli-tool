package main

import "fmt"

func commandHelp(c *config) error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	commands := getCommands()
	for command := range commands {
		fmt.Printf("%s: %s\n", command, commands[command].description)
	}

	return nil
}
