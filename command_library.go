package main

import (
	"errors"
	"fmt"
)

type cliCommand struct {
	name        string
	description string
	helptext    string
	callback    func(cfg *config, params ...string) error
}

type commandLibrary map[string]cliCommand

func (l *commandLibrary) PrintCommands() {
	fmt.Println("Available Commands")
	fmt.Println("------------------")
	for _, entry := range *l {
		fmt.Printf(" - %s: %s\n", entry.name, entry.description)

	}
}

func (l commandLibrary) getCommand(name string) (cliCommand, error) {
	command, ok := l[name]
	if !ok {
		return command, errors.New("no such command")
	}

	return command, nil
}

func NewCommandLibrary() *commandLibrary {
	return &commandLibrary{
		"help": {
			name:        "help",
			description: "Displays a help message",
			helptext:    "Use: 'help <command>'",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			helptext:    "Use: 'exit' to exit the program.",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays 20 location areas in the Pokemon World. Each subsequent call to map will display then next 20 locations",
			helptext:    "Use: 'map' to show 20 areas of the Pokemon World.",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 location areas.",
			helptext:    "Use: 'mapb' to show previous 20 areas of the Pokemon World.",
			callback:    commandMapB,
		},
		"explore": {
			name:        "explore",
			description: "Displays pokemon for the given area.",
			helptext:    "Use: 'explore <location-area>' to display pokemon for the <location-area>",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempt to catch the given pokemon",
			helptext:    "Use: 'catch <pokemon>' to attempt to catch <pokemon>",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Review data from caught pokemon",
			helptext:    "Use: 'inspect <pokemon>' to review data of caught pokemon",
			callback:    commandInspect,
		},

		"pokedex": {
			name:        "pokedex",
			description: "Lists all pokemon that you have caught",
			helptext:    "Use: 'pokedex' print all caught pokemon",
			callback:    commandPokedex,
		},
	}
}
