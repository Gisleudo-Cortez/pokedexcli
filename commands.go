package main

import (
	"fmt"
	"os"
)

// Create commands for usage in cli
func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}

	return nil
}

func commandExit() error {
	os.Exit(0)
	return nil
}
