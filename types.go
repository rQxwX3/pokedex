package main

import (
	"github.com/rQxwX3/pokedex/internal/api"
	"github.com/rQxwX3/pokedex/internal/pokecache"
)

type cliCmdsMap map[string]cliCmd
type pokedex map[string]api.Pokemon

type cliCmd struct {
	name        string
	description string
	callback    func(conf *config) error
	argsCount   int
}

type config struct {
	cmdsMap cliCmdsMap
	pokedex pokedex
	args    []string
	next    string
	prev    string
	cache   *pokecache.Cache
}
