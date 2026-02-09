package main

import (
	"fmt"
	"os"
)

func CommandExit(conf *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)

	return nil
}

func CommandHelp(conf *config) error {
	fmt.Println("Welcome to the Pokedex!\nUsage:\n")

	for _, cmd := range nameToCommand {
		fmt.Println(cmd.name + ": " + cmd.description)
	}

	return nil
}

func CommandMap(conf *config) error {
	url := "https://pokeapi.co/api/v2/location-area"

	if conf.Next != "" {
		url = conf.Next
	}

	var locations Locations
	if err := APIGet(url, conf.Cache, &locations); err != nil {
		return err
	}

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	conf.Next = locations.Next
	conf.Prev = locations.Previous

	return nil
}

func CommandMapBack(conf *config) error {
	if conf.Prev == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	var locations Locations
	if err := APIGet(conf.Prev, conf.Cache, &locations); err != nil {
		return err
	}

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	conf.Next = locations.Next
	conf.Prev = locations.Previous

	return nil
}
