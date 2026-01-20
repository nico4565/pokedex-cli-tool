package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/mitchellh/go-wordwrap"
)

func commandHelp(c *config, args ...string) error {
	fmt.Println("Welcome to the Pokedex help menu!")

	commands := getCommands()

	//Collect and sort command names
	keys := make([]string, 0, len(commands))
	//maxLen to dinamically set column width
	maxLen := 0
	for k := range commands {
		keys = append(keys, k)
		if len(k) > maxLen {
			maxLen = len(k)
		}
	}
	sort.Strings(keys)

	//Finally print aligned output using world wrapper library
	for _, k := range keys {
		wrapped := wordwrap.WrapString(commands[k].description, 80)
		wrapped = strings.Replace(wrapped, "\n", "\n"+fmt.Sprintf("%*s", maxLen+3, ""), -1)

		fmt.Printf("%-*s : %s\n", maxLen, k, wrapped)
	}

	return nil
}
