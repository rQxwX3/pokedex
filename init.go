package main

import (
	"github.com/rQxwX3/pokedex/internal/pokecache"
	"time"
)

func initConfig() config {
	return config{
		pokedex: pokedex{},
		cmdsMap: initCommandsMap(),
		args:    []string{},
		next:    "",
		prev:    "",
		cache:   pokecache.NewCache(5 * time.Second),
	}
}

func initCommandsMap() cliCmdsMap {
	var cliCommandsMap = cliCmdsMap{}

	cliCommandsMap["exit"] = cliCmd{
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    CommandExit,
		argsCount:   0,
	}

	cliCommandsMap["help"] = cliCmd{
		name:        "help",
		description: "Print Pokedex help message",
		callback:    CommandHelp,
		argsCount:   0,
	}

	cliCommandsMap["map"] = cliCmd{
		name:        "map",
		description: "Print next 20 locations",
		callback:    CommandMap,
		argsCount:   0,
	}

	cliCommandsMap["mapb"] = cliCmd{
		name:        "mapb",
		description: "Print previous 20 locations",
		callback:    CommandMapBack,
		argsCount:   0,
	}

	cliCommandsMap["explore"] = cliCmd{
		name:        "explore [location]",
		description: "Print all Pokemons found in specified location",
		callback:    CommandExplore,
		argsCount:   1,
	}

	cliCommandsMap["catch"] = cliCmd{
		name:        "catch [pokemon]",
		description: "Attempt to catch specified Pokemon",
		callback:    CommandCatch,
		argsCount:   1,
	}

	cliCommandsMap["inspect"] = cliCmd{
		name:        "inspect [pokemon]",
		description: "Get info about specified Pokemon",
		callback:    CommandInspect,
		argsCount:   1,
	}

	return cliCommandsMap
}
