package pokeapi

import (
	"net/http"
	"time"
)

// struct that contains client, httpClient is unexported, only accessible in the same package, allows for
// more control and safety
type Client struct {
	httpClient http.Client
}

func NewClient(timeout time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
