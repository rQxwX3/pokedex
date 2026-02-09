package main

import (
	"encoding/json"
	"errors"
	"github.com/rQxwX3/pokedex/internal/pokecache"
	"io"
	"net/http"
)

func APIGet(url string, cache *pokecache.Cache, storage any) error {
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
		return errors.New("API request was unsuccessful")
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
