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

// commandMapWrapper: a wrapper to fit the commandMap into the expected callback signature
func commandMapWrapper() func() error {
	return func() error {
		return commandMap(&config)
	}
}

// commandMapbWrapper: a wrapper to fit the commandMapb into the expected callback signature
func commandMapbWrapper() func() error {
	return func() error {
		return commandMapb(&config)
	}
}

// commandMap
// displays the names of 20 location areas in the Pokemon world. Each subsequent call to map should display the next 20 locations

var config = configMaps{
	NextURL:     "https://pokeapi.co/api/v2/location-area/?limit=20",
	PreviousURL: "",
}

var offset = 0

func commandMap(config *configMaps) error {
	if config.NextURL == "" {
		return fmt.Errorf("No URL provided")
	}
	offset += 20
	config.PreviousURL = config.NextURL
	next := fmt.Sprintf("&offset=%v", offset)
	config.NextURL = fmt.Sprintf("https://pokeapi.co/api/v2/location-area/?limit=20%s", next)

	names, err := getLocationAreaURL(config.NextURL)
	if err != nil {
		return err
	}
	for _, nm := range names {
		fmt.Println(nm)
	}
	return nil
}

// commandMapb
// instead of displaying the next 20 locations, it displays the previous 20 locations.
func commandMapb(config *configMaps) error {
	if config.NextURL == "" {
		return fmt.Errorf("No URL provided")
	}
	if offset == 0 {

	}
	offset -= 20
	config.PreviousURL = config.NextURL
	next := fmt.Sprintf("&offset=%v", offset)
	config.NextURL = fmt.Sprintf("https://pokeapi.co/api/v2/location-area/?limit=20%s", next)

	names, err := getLocationAreaURL(config.NextURL)
	if err != nil {
		return err
	}
	for _, nm := range names {
		fmt.Println(nm)
	}
	return nil
}

func test() error {
	fmt.Println(config)
	return nil
}
