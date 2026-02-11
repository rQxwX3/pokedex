package types

import (
	"github.com/rQxwX3/pokedex/internal/api"
	"github.com/rQxwX3/pokedex/internal/pokecache"
)

type CliCmdsMap map[string]CliCmd
type Pokedex map[string]api.Pokemon

type CliCmd struct {
	Name        string
	Description string
	Callback    func(conf *Config) error
	ArgsCount   int
}

type Config struct {
	CmdsMap CliCmdsMap
	Pokedex Pokedex
	Args    []string
	Next    string
	Prev    string
	Cache   *pokecache.Cache
}
