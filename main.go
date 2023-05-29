package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

type commandLibrary map[string]cliCommand

func (c commandLibrary) getCommand(name string) (cliCommand, error) {
	command, ok := c[name]
	if !ok {
		return command, errors.New("no such command")
	}

	return command, nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	return nil
}

func commandExit() error {
	os.Exit(0)
	return nil
}

func commandMap() error {
	pokeapi.GetAreas()
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
	}

	for {
		fmt.Printf("pokedex > ")
		input, _ := reader.ReadString('\n')
		input = strings.Replace(input, "\n", "", -1)

		command, err := commands.getCommand(input)
		if err != nil {
			fmt.Printf("error: %s\n", err)
		} else {
			command.callback()
		}

	}

}
