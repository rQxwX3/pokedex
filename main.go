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
	Next  string
	Prev  string
	Cache *pokecache.Cache
}

type cliCommand struct {
	name        string
	description string
	callback    func(conf *config) error
}

var nameToCommand = map[string]cliCommand{}

func initCommandMap() {
	nameToCommand["exit"] = cliCommand{
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    CommandExit,
	}

	nameToCommand["help"] = cliCommand{
		name:        "help",
		description: "Print Pokedex help message",
		callback:    CommandHelp,
	}

	nameToCommand["map"] = cliCommand{
		name:        "map",
		description: "Print next 20 locations",
		callback:    CommandMap,
	}

	nameToCommand["mapb"] = cliCommand{
		name:        "mapb",
		description: "Print previous 20 locations",
		callback:    CommandMapBack,
	}
}

func main() {
	initCommandMap()

	conf := config{
		Next:  "",
		Prev:  "",
		Cache: pokecache.NewCache(5 * time.Second),
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

		err := cliCommandStruct.callback(&conf)
		if err != nil {
			fmt.Printf("Command callback failed: %v\n", err)
			os.Exit(1)
		}
	}
}
