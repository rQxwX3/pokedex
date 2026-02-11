package api

import (
	"encoding/json"
	"errors"
	"github.com/rQxwX3/pokedex/internal/pokecache"
	"io"
	"net/http"
)

type Locations struct {
	Next     string
	Previous string
	Results  []Location
}

type Location struct {
	Name string
}

type LocationArea struct {
	Pokemons []struct {
		Pokemon struct {
			Name string `json:"name"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

type Pokemon struct {
	Experience int `json:"base_experience"`
}

func Get(url string, cache *pokecache.Cache, storage any) error {
	if cachedData, ok := cache.Get(url); ok {
		if err := json.Unmarshal(cachedData, storage); err != nil {
			return err
		}

		return nil
	}

	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return errors.New("API request failed")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(body, storage); err != nil {
		return err
	}

	cache.Add(url, body)

	return nil
}
