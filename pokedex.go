package main

import (
	"fmt"

	"github.com/localman-src/pokedex/internal/pokeapi"
	"github.com/localman-src/pokedex/internal/structs"
)

type PokedexEntry struct {
	Name        string
	Count       int
	Height      int
	Weight      int
	Stats       []structs.PokemonStat
	Types       []structs.PokemonType
	Description []structs.FlavorText
}
type Pokedex map[string]PokedexEntry

func (p Pokedex) Add(e PokedexEntry) {
	entry, ok := p[e.Name]
	if !ok {
		p[e.Name] = e
	} else {
		entry.Count++
		p[e.Name] = entry
	}
}

func (p Pokedex) Get(name string) (entry PokedexEntry, err error) {
	entry, ok := p[name]
	if !ok {
		return entry, fmt.Errorf("pokemon not caught yet: %s", name)
	}

	return entry, err
}

func NewPokedex() Pokedex {
	return Pokedex{}
}

func NewPokedexEntry(pokemon *structs.Pokemon) (PokedexEntry, error) {

	species, err := pokeapi.CachedGet(pokemon.Species.ToAPIResource())
	if err != nil {
		return PokedexEntry{}, err
	}
	return PokedexEntry{
		Name:        pokemon.Name,
		Count:       1,
		Height:      pokemon.Height,
		Weight:      pokemon.Weight,
		Stats:       pokemon.Stats,
		Types:       pokemon.Types,
		Description: species.FlavorTextEntries,
	}, nil
}
