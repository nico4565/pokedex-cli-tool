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
		fmt.Printf("Your command was: %v\n", parsed[0])

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
