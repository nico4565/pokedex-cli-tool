package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func StartRepl() {
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
			fmt.Print("Unknown command")
		}

		err := command.callback()
		if err != nil {
			fmt.Print(err)
		}
	}
}

func commandExit() error {
	fmt.Print("Closing the Pokedex... Goodbye!")
	os.Exit(0)

	return nil
}

func commandHelp() error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	commands := getCommands()
	for command := range commands {
		fmt.Printf("%s: %s\n", command, commands[command].description)
	}

	return nil
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

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the pokedex",
			callback:    commandExit,
		}, "help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
	}
}
