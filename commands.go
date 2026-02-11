package main

import (
	"fmt"
	"github.com/rQxwX3/pokedex/internal/api"
	"github.com/rQxwX3/pokedex/internal/types"
	"math/rand"
	"os"
)

func CommandExit(conf *types.Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)

	return nil
}

func CommandHelp(conf *types.Config) error {
	fmt.Println("Welcome to the Pokedex!\nUsage:")
	fmt.Println()

	for _, cmd := range conf.CmdsMap {
		fmt.Println(cmd.Name + ": " + cmd.Description)
	}

	return nil
}

func CommandMap(conf *types.Config) error {
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

func CommandMapBack(conf *types.Config) error {
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

func CommandExplore(conf *types.Config) error {
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
	xp := pokemon.Experience + 1

	return (10 + 1000/xp) % 100
}

func CommandCatch(conf *types.Config) error {
	pokemonName := conf.Args[0]

	url := "https://pokeapi.co/api/v2/pokemon/" + pokemonName

	var pokemon api.Pokemon
	if err := api.Get(url, conf.Cache, &pokemon); err != nil {
		return err
	}

	pokemon.Name = pokemonName

	fmt.Println("Throwing a Pokeball at " + pokemonName + "...")
	pokedex := conf.Pokedex

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

func CommandInspect(conf *types.Config) error {
	pokemon, ok := conf.Pokedex[conf.Args[0]]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	fmt.Println("Name:", pokemon.Name)
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

func CommandPokedex(conf *types.Config) error {
	fmt.Println("Your Pokedex:")

	for _, pokemon := range conf.Pokedex {
		fmt.Println("  -", pokemon.Name)
	}

	return nil
}
