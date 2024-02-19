package pokeapi

import (
	"net/http"
	"time"
	"github.com/hackastak/pokedexcli/internal/pokecache"

)

type Client struct {
	cache pokecache.Cache
	httpClient http.Client
}

func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
