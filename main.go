package main

import (
	"bufio"
	"fmt"
	"github.com/rQxwX3/pokedex/internal/pokecache"
	"os"
	"strings"
	"time"
)

func CleanInput(text string) []string {
	words := strings.Fields(text)

	for i := range words {
		words[i] = strings.ToLower(words[i])
	}

	return words
}

type config struct {
	Next    string
	Prev    string
	Cache   *pokecache.Cache
	Args    []string
	Pokedex pokedex
}

type cliCommand struct {
	name           string
	description    string
	callback       func(conf *config) error
	argumentsCount int
}

type pokedex struct {
	pokemons []string
}

var nameToCommand = map[string]cliCommand{}

func initCommandMap() {
	nameToCommand["exit"] = cliCommand{
		name:           "exit",
		description:    "Exit the Pokedex",
		callback:       CommandExit,
		argumentsCount: 0,
	}

	nameToCommand["help"] = cliCommand{
		name:           "help",
		description:    "Print Pokedex help message",
		callback:       CommandHelp,
		argumentsCount: 0,
	}

	nameToCommand["map"] = cliCommand{
		name:           "map",
		description:    "Print next 20 locations",
		callback:       CommandMap,
		argumentsCount: 0,
	}

	nameToCommand["mapb"] = cliCommand{
		name:           "mapb",
		description:    "Print previous 20 locations",
		callback:       CommandMapBack,
		argumentsCount: 0,
	}

	nameToCommand["explore"] = cliCommand{
		name:           "explore [location]",
		description:    "Print all Pokemons found in specified location",
		callback:       CommandExplore,
		argumentsCount: 1,
	}

	nameToCommand["catch"] = cliCommand{
		name:           "catch [pokemon]",
		description:    "Attempt to catch specified Pokemon",
		callback:       CommandCatch,
		argumentsCount: 1,
	}

	nameToCommand["inspect"] = cliCommand{
		name:           "inspect [pokemon]",
		description:    "Get info about specified Pokemon",
		callback:       CommandCatch,
		argumentsCount: 1,
	}
}

func main() {
	initCommandMap()

	conf := config{
		Next:  "",
		Prev:  "",
		Cache: pokecache.NewCache(5 * time.Second),
		Args:  []string{},
	}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")

		scanner.Scan()
		words := CleanInput(scanner.Text())

		cliCommandStruct, ok := nameToCommand[words[0]]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		conf.Args = append(conf.Args, words[1:]...)

		err := cliCommandStruct.callback(&conf)
		if err != nil {
			fmt.Printf("Command callback failed: %v\n", err)
			os.Exit(1)
		}
	}
}
