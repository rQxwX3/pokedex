package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func CleanInput(text string) []string {
	words := strings.Fields(text)

	for i := 0; i < len(words); i++ {
		words[i] = strings.ToLower(words[i])
	}

	return words
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		words := CleanInput(input)
		fmt.Println("Your command was:", words[0])
	}
}
