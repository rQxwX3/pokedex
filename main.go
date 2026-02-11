package main

import (
	"github.com/rQxwX3/pokedex/internal/repl"
)

func main() {
	conf := initConfig()

	repl.Run(&conf)
}
