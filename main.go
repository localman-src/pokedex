package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type config struct {
	MapOffset     int
	ResourceLimit int
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	commands := NewCommandLibrary()
	config := config{
		MapOffset:     0,
		ResourceLimit: 20,
	}

	for {
		fmt.Printf("pokedex > ")
		input, _ := reader.ReadString('\n')
		input = strings.Replace(input, "\n", "", -1)
		params := strings.Split(input, " ")
		fmt.Println(params)

		command, err := commands.getCommand(params[0])
		if err != nil {
			fmt.Printf("error: %s\n", err)
		} else {
			if len(params) == 1 {
				command.callback(&config)
			} else {
				command.callback(&config, params[1:]...)
			}

		}

	}

}
