package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/localman-src/pokedex/internal/pokeapi"
)

type config struct {
	NextMap string
	PrevMap string
}

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config) error
}

type commandLibrary map[string]cliCommand

func (c commandLibrary) getCommand(name string) (cliCommand, error) {
	command, ok := c[name]
	if !ok {
		return command, errors.New("no such command")
	}

	return command, nil
}

func commandHelp(cfg *config) error {
	fmt.Println("Welcome to the Pokedex!")
	return nil
}

func commandExit(cfg *config) error {
	os.Exit(0)
	return nil
}

func commandMap(cfg *config) error {

	resp, err := pokeapi.GetAreas(cfg.NextMap)
	if err != nil {
		fmt.Printf("error: %s\n", err)
	} else {
		pokeapi.PrintResourceList(*resp)
		cfg.NextMap = resp.Next
		cfg.PrevMap = resp.Previous
	}

	return nil
}

func commandMapB(cfg *config) error {
	if cfg.PrevMap != "" {
		resp, err := pokeapi.GetAreas(cfg.PrevMap)
		if err != nil {
			fmt.Printf("error: %s\n", err)
		} else {
			pokeapi.PrintResourceList(*resp)
			cfg.NextMap = resp.Next
			cfg.PrevMap = resp.Previous
		}
	}
	return nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	commands := commandLibrary{
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
	}

	config := config{
		NextMap: "https://pokeapi.co/api/v2/location-area/",
		PrevMap: "",
	}

	for {
		fmt.Printf("pokedex > ")
		input, _ := reader.ReadString('\n')
		input = strings.Replace(input, "\n", "", -1)

		command, err := commands.getCommand(input)
		if err != nil {
			fmt.Printf("error: %s\n", err)
		} else {
			command.callback(&config)
		}

	}

}
