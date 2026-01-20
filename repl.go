package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/nico4565/pokedex-cli-tool/internal/pokeapi"
)

type config struct {
	httpClient    pokeapi.Client
	nextUrl       *string
	previousUrl   *string
	caughtPokemon map[string]pokeapi.Pokemon
}

type cliCommand struct {
	name        string
	description string
	callback    func(c *config, args ...string) error
}

func StartRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		text := scanner.Text()
		if len(text) == 0 {
			continue
		}

		parsed := cleanInput(text)
		command, exists := getCommands()[parsed[0]]
		if !exists {
			fmt.Println("Unknown command")
			continue
		}

		args := []string{}

		if len(parsed) > 1 {
			args = parsed[1:]
		}

		err := command.callback(cfg, args...)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Prints the next page of location areas, 20 at a time",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Prints the previous page of location areas, 20 at a time, if there is no previous page it will just print a message",
			callback:    commandMapBack,
		},
		"explore": {
			name:        "explore",
			description: "Takes a location area after the explore command and gives back a list of pokemon present in the area. If you don't specify a location are the command will print a message asking to add one.",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Takes a pokemon (name or id) after the catch command, it catch can either be succesfull or not, if it is you get to save the pokemon on your pokedex. If you don't specify a pokemon the command will print a message asking to add one.",
			callback:    commandCatch,
		},
	}
}

func cleanInput(text string) []string {
	res := []string{}
	text = strings.ToLower(text)
	text = strings.Trim(text, " ")
	i := strings.Index(text, " ")

	for i > -1 {
		res = append(res, text[:i])
		text = text[i+1:]
		if text == " " {
			return res
		}
		i = strings.Index(text, " ")
	}

	return append(res, text)
}
