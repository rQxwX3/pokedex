package main

import (
	"fmt"
	"github.com/rQxwX3/pokedex/internal/api"
	"math/rand"
	"os"
)

func CommandExit(conf *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)

	return nil
}

func CommandHelp(conf *config) error {
	fmt.Println("Welcome to the Pokedex!\nUsage:")
	fmt.Println()

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

	var locations api.Locations
	if err := api.Get(url, conf.Cache, &locations); err != nil {
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

	var locations api.Locations
	if err := api.Get(conf.Prev, conf.Cache, &locations); err != nil {
		return err
	}

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	conf.Next = locations.Next
	conf.Prev = locations.Previous

	return nil
}

func CommandExplore(conf *config) error {
	url := "https://pokeapi.co/api/v2/location-area/" + conf.Args[0]

	var locationArea api.LocationArea
	if err := api.Get(url, conf.Cache, &locationArea); err != nil {
		return err
	}

	for _, pokemon := range locationArea.Pokemons {
		fmt.Println(pokemon.Pokemon.Name)
	}

	return nil
}

func calculateChance(pokemon api.Pokemon) int {
	return (1000 / pokemon.Experience) % 100
}

func CommandCatch(conf *config) error {
	pokemonName := conf.Args[0]

	url := "https://pokeapi.co/api/v2/pokemon/" + pokemonName

	var pokemon api.Pokemon
	if err := api.Get(url, conf.Cache, &pokemon); err != nil {
		return err
	}

	fmt.Println("Throwing a Pokeball at " + pokemonName + "...")
	pokedex := conf.Pokedex

	if rand.Intn(101) < calculateChance(pokemon) {
		pokedex.pokemons = append(pokedex.pokemons, pokemonName)
		fmt.Println(pokemonName + " was caught!")
	} else {
		fmt.Println(pokemonName + " escaped!")
	}

	return nil
}

func CommandInspect(conf *Config) error {
	pokedex := 	
}
