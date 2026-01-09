package main

import (
	"time"

	"github.com/nico4565/pokedex-cli-tool/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(4 * time.Second)
	cfg := &config{
		httpClient: pokeClient,
	}

	StartRepl(cfg)
}
