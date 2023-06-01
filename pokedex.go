package main

import (
	"encoding/gob"
	"fmt"
	"os"

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
	Description []structs.FlavorText `json:"-"`
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

func (p Pokedex) Save() error {
	f, err := os.Create("./pokedex.sav")
	if err != nil {
		return err
	}
	defer f.Close()

	enc := gob.NewEncoder(f)
	err = enc.Encode(p)
	if err != nil {
		return err
	}

	return nil
}

func (p *Pokedex) Load() (*Pokedex, error) {
	var loadedDex Pokedex

	data, err := os.Open("./pokedex.sav")
	if err != nil {
		return &loadedDex, err
	}
	defer data.Close()

	enc := gob.NewDecoder(data)
	err = enc.Decode(&loadedDex)
	if err != nil {
		return &loadedDex, err
	}

	return &loadedDex, nil
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
