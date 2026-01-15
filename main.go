package main

import (
	"time"

	"github.com/nico4565/pokedex-cli-tool/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	cfg := &config{
		httpClient: pokeClient,
	}

	StartRepl(cfg)
}
