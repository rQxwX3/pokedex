package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func CleanInput(text string) []string {
	words := strings.Fields(text)

	for i := range words {
		words[i] = strings.ToLower(words[i])
	}

	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var commandMap = map[string]cliCommand{}

func CommandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)

	return nil
}

func CommandHelp() error {
	fmt.Println("Welcome to the Pokedex!\nUsage:\n")

	for _, cmd := range commandMap {
		fmt.Println(cmd.name + ": " + cmd.description)
	}

	return nil
}

func initCommandMap() {
	commandMap["exit"] = cliCommand{
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    CommandExit,
	}

	commandMap["help"] = cliCommand{
		name:        "help",
		description: "Print Pokedex help message",
		callback:    CommandHelp,
	}
}

func main() {
	initCommandMap()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")

		scanner.Scan()
		words := CleanInput(scanner.Text())

		cliCommandStruct, ok := commandMap[words[0]]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		err := cliCommandStruct.callback()
		if err != nil {
			fmt.Printf("Command callback failed: %v\n", err)
			os.Exit(1)
		}
	}
}
