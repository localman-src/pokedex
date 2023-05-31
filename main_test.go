package main

import "testing"

func TestCommands(t *testing.T) {
	library := NewCommandLibrary()

	for _, command := range *library {
		_, err := library.getCommand(command.name)
		if err != nil {
			t.Errorf("expected to find command: %s", command.name)
		}
	}

	_, err := library.getCommand("THIS IS NOT A REAL COMMAND")
	if err != nil {
		return
	}
	t.Errorf("expected not to find command")
}
