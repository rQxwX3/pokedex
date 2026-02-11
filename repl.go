package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	prompt       = "Pokedex > "
	unknowncmd   = "Unknown command:"
	noarg        = "This command requires argument(s). Run help command for more info."
	callbackfail = "Command callback failed: "
)

func repl(conf *config) error {
	scanner := bufio.NewScanner(os.Stdin)
	commandsMap := conf.cmdsMap

	for {
		fmt.Print(prompt)

		scanner.Scan()
		words := cleanInput(scanner.Text())

		cmd, args, ok := getCliCommand(words, &commandsMap)
		if !ok {
			continue
		}

		conf.args = append(conf.args, args...)

		if err := cmd.callback(conf); err != nil {
			fmt.Println(callbackfail, err)
		}

		conf.args = conf.args[:0]
	}
}

func getCliCommand(words []string,
	commandsMap *cliCmdsMap) (cliCmd, []string, bool) {
	cmd, ok := (*commandsMap)[words[0]]

	if !ok {
		fmt.Println(unknowncmd, words[0])
		return cliCmd{}, nil, false
	}

	args := words[1:]
	if len(args) != cmd.argsCount {
		fmt.Println(noarg)
		return cliCmd{}, nil, false
	}

	return cmd, args, true
}

func cleanInput(text string) []string {
	words := strings.Fields(text)

	for i := range words {
		words[i] = strings.ToLower(words[i])
	}

	return words
}
