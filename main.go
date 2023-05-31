package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

type config struct {
	MapOffset       int
	ResourceLimit   int
	CurrentLocation string
	Commands        *commandLibrary
	Pokedex         Pokedex
	prng            *rand.Rand
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	config := config{
		MapOffset:       0,
		ResourceLimit:   20,
		CurrentLocation: "canalave-city-area",
		Commands:        NewCommandLibrary(),
		Pokedex:         NewPokedex(),
		prng:            rand.New(rand.NewSource(90)),
	}

	for {
		fmt.Printf("pokedex > ")
		input, _ := reader.ReadString('\n')
		input = strings.Replace(input, "\n", "", -1)
		params := strings.Split(input, " ")
		//fmt.Println(params)

		command, err := config.Commands.getCommand(params[0])
		if err != nil {
			fmt.Printf("error: %s\n", err)
			continue
		}

		if len(params) == 1 {
			err = command.callback(&config)
			if err != nil {
				fmt.Printf("error: %s\n", err)
			}
			continue
		}

		err = command.callback(&config, params[1:]...)
		if err != nil {
			fmt.Printf("error: %s\n", err)
		}

	}

}
