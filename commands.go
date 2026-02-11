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

	for _, cmd := range conf.cmdsMap {
		fmt.Println(cmd.name + ": " + cmd.description)
	}

	return nil
}

func CommandMap(conf *config) error {
	url := "https://pokeapi.co/api/v2/location-area"

	if conf.next != "" {
		url = conf.next
	}

	var locations api.Locations
	if err := api.Get(url, conf.cache, &locations); err != nil {
		return err
	}

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	conf.next = locations.Next
	conf.prev = locations.Previous

	return nil
}

func CommandMapBack(conf *config) error {
	if conf.prev == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	var locations api.Locations
	if err := api.Get(conf.prev, conf.cache, &locations); err != nil {
		return err
	}

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	conf.next = locations.Next
	conf.prev = locations.Previous

	return nil
}

func CommandExplore(conf *config) error {
	url := "https://pokeapi.co/api/v2/location-area/" + conf.args[0]

	var locationArea api.LocationArea
	if err := api.Get(url, conf.cache, &locationArea); err != nil {
		return err
	}

	for _, pokemon := range locationArea.Pokemons {
		fmt.Println(pokemon.Pokemon.Name)
	}

	return nil
}

func calculateChance(pokemon api.Pokemon) int {
	xp := pokemon.Experience + 1

	return (10 + 1000/xp) % 100
}

func CommandCatch(conf *config) error {
	pokemonName := conf.args[0]

	url := "https://pokeapi.co/api/v2/pokemon/" + pokemonName

	var pokemon api.Pokemon
	if err := api.Get(url, conf.cache, &pokemon); err != nil {
		return err
	}

	fmt.Println("Throwing a Pokeball at " + pokemonName + "...")
	pokedex := conf.pokedex

	if rand.Intn(101) < calculateChance(pokemon) {
		if _, ok := pokedex[pokemonName]; !ok {
			pokedex[pokemonName] = pokemon
		}

		fmt.Println(pokemonName + " was caught!")
	} else {
		fmt.Println(pokemonName + " escaped!")
	}

	return nil
}

func CommandInspect(conf *config) error {
	pokemon, ok := conf.pokedex[conf.args[0]]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	fmt.Println("Name:", conf.args[0])
	fmt.Println("Height:", pokemon.Height)
	fmt.Println("Weight:", pokemon.Weight)

	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Println("  -", stat.StatInfo.Name+":", stat.Value)
	}

	fmt.Println("Types:")
	for _, pokeType := range pokemon.Types {
		fmt.Println("  -", pokeType.Type.Name)
	}

	return nil
}
