package main

//"errors"
import (
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	fmt.Println("Your Pokedex:")

	for _, name := range cfg.caughtPokemon {
		fmt.Printf(" - %s\n", name.Name)
	}

	return nil
}
