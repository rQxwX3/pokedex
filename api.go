package main

import (
	"encoding/json"
	"errors"
	"net/http"
)

func APIGet(url string, storage any) error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)

	if err = decoder.Decode(storage); err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return errors.New("API request was unsuccessful")
	}

	return nil
}
