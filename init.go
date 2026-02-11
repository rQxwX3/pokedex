package main

import (
	"github.com/rQxwX3/pokedex/internal/pokecache"
	"github.com/rQxwX3/pokedex/internal/types"
	"time"
)

func initConfig() types.Config {
	return types.Config{
		Pokedex: types.Pokedex{},
		CmdsMap: initCommandsMap(),
		Args:    []string{},
		Next:    "",
		Prev:    "",
		Cache:   pokecache.NewCache(5 * time.Second),
	}
}

func initCommandsMap() types.CliCmdsMap {
	var cliCommandsMap = types.CliCmdsMap{}

	cliCommandsMap["exit"] = types.CliCmd{
		Name:        "exit",
		Description: "Exit the Pokedex",
		Callback:    CommandExit,
		ArgsCount:   0,
	}

	cliCommandsMap["help"] = types.CliCmd{
		Name:        "help",
		Description: "Print Pokedex help message",
		Callback:    CommandHelp,
		ArgsCount:   0,
	}

	cliCommandsMap["map"] = types.CliCmd{
		Name:        "map",
		Description: "Print next 20 locations",
		Callback:    CommandMap,
		ArgsCount:   0,
	}

	cliCommandsMap["mapb"] = types.CliCmd{
		Name:        "mapb",
		Description: "Print previous 20 locations",
		Callback:    CommandMapBack,
		ArgsCount:   0,
	}

	cliCommandsMap["explore"] = types.CliCmd{
		Name:        "explore [location]",
		Description: "Print all Pokemons found in specified location",
		Callback:    CommandExplore,
		ArgsCount:   1,
	}

	cliCommandsMap["catch"] = types.CliCmd{
		Name:        "catch [pokemon]",
		Description: "Attempt to catch specified Pokemon",
		Callback:    CommandCatch,
		ArgsCount:   1,
	}

	cliCommandsMap["inspect"] = types.CliCmd{
		Name:        "inspect [pokemon]",
		Description: "Get info about specified Pokemon",
		Callback:    CommandInspect,
		ArgsCount:   1,
	}

	cliCommandsMap["pokedex"] = types.CliCmd{
		Name:        "pokedex",
		Description: "List caught pokemons",
		Callback:    CommandPokedex,
		ArgsCount:   0,
	}

	return cliCommandsMap
}
