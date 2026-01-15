package pokeapi

import (
	"net/http"
	"time"

	"github.com/nico4565/pokedex-cli-tool/internal/pokecache"
)

// struct that contains client, httpClient is unexported, only accessible in the same package, allows for
// more control and safety
type Client struct {
	httpClient http.Client
	cache      pokecache.Cache
}

func NewClient(timeout, interval time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		cache: *pokecache.NewCache(interval),
	}
}
