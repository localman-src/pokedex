package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/localman-src/pokedex/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config, params ...string) error
}

type commandLibrary map[string]cliCommand

func (c commandLibrary) getCommand(name string) (cliCommand, error) {
	command, ok := c[name]
	if !ok {
		return command, errors.New("no such command")
	}

	return command, nil
}

func commandHelp(cfg *config, params ...string) error {
	fmt.Println("Welcome to the Pokedex!")
	return nil
}

func commandExit(cfg *config, params ...string) error {
	os.Exit(0)
	return nil
}

func commandMap(cfg *config, params ...string) error {

	resp, err := pokeapi.GetResourceList(pokeapi.EndPoints.LocationArea, cfg.MapOffset, cfg.ResourceLimit)
	if err != nil {
		fmt.Printf("error: %s\n", err)
	} else {
		resp.Print()
		cfg.MapOffset += cfg.ResourceLimit
		fmt.Printf("New Offset: %d\n", cfg.MapOffset)
	}

	return nil
}

func commandMapB(cfg *config, params ...string) error {

	newOffset := cfg.MapOffset - cfg.ResourceLimit
	if newOffset <= 0 {
		newOffset = 0
	}
	resourceList, err := pokeapi.GetResourceList("/location-area", newOffset, cfg.ResourceLimit)
	if err != nil {
		fmt.Printf("error: %s\n", err)
	} else {
		resourceList.Print()
		cfg.MapOffset = newOffset
		fmt.Printf("New Offset: %d\n", cfg.MapOffset)
	}
	return nil
}

func commandExplore(cfg *config, params ...string) error {
	locationArea, err := pokeapi.GetLocationArea(params[0])
	if err != nil {
		fmt.Printf("error: %s\n", err)
	} else {
		locationArea.PrintPokemonEncounters()
	}

	return nil
}

func NewCommandLibrary() commandLibrary {
	return commandLibrary{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays 20 location areas in the Pokemon World.\n Each subsequent call to map will display then next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 location areas.",
			callback:    commandMapB,
		},
		"explore": {
			name:        "explore",
			description: "Use: explore <area>. Displays pokemon for the given area.",
			callback:    commandExplore,
		},
	}
}
