package repl

import (
	"bufio"
	"fmt"
	"github.com/rQxwX3/pokedex/internal/types"
	"os"
	"strings"
)

const (
	prompt       = "Pokedex > "
	unknowncmd   = "Unknown command:"
	noarg        = "This command requires argument(s). Run help command for more info."
	callbackfail = "Command callback failed: "
)

func Run(conf *types.Config) error {
	scanner := bufio.NewScanner(os.Stdin)
	commandsMap := conf.CmdsMap

	for {
		fmt.Print(prompt)

		scanner.Scan()
		words := cleanInput(scanner.Text())

		cmd, args, ok := getCliCommand(words, &commandsMap)
		if !ok {
			continue
		}

		conf.Args = append(conf.Args, args...)

		if err := cmd.Callback(conf); err != nil {
			fmt.Println(callbackfail, err)
		}

		conf.Args = conf.Args[:0]
	}
}

func getCliCommand(words []string,
	commandsMap *types.CliCmdsMap) (types.CliCmd, []string, bool) {
	if len(words) < 1 {
		return types.CliCmd{}, nil, false
	}

	cmd, ok := (*commandsMap)[words[0]]

	if !ok {
		fmt.Println(unknowncmd, words[0])
		return types.CliCmd{}, nil, false
	}

	args := words[1:]
	if len(args) != cmd.ArgsCount {
		fmt.Println(noarg)
		return types.CliCmd{}, nil, false
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
