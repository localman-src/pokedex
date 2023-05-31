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
	MapOffset     int
	ResourceLimit int
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

	resp, err := pokeapi.GetResourceList(pokeapi.EndPoints.LocationArea, cfg.MapOffset, cfg.ResourceLimit)
	if err != nil {
		fmt.Printf("error: %s\n", err)
	} else {
		resp.Print()
		cfg.MapOffset += 20
		fmt.Printf("New Offset: %d\n", cfg.MapOffset)
	}

	return nil
}

func commandMapB(cfg *config) error {

	newOffset := cfg.MapOffset - cfg.ResourceLimit
	if newOffset <= 0 {
		newOffset = 0
	}
	resp, err := pokeapi.GetResourceList("/location-area", newOffset, cfg.ResourceLimit)
	if err != nil {
		fmt.Printf("error: %s\n", err)
	} else {
		resp.Print()
		cfg.MapOffset = newOffset
		fmt.Printf("New Offset: %d\n", cfg.MapOffset)
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
		MapOffset:     0,
		ResourceLimit: 20,
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
