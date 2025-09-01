package main

import (
	"fmt"
	"os"
	"pokedexcli/internal/config"
)

func commandExit(cfg *config.Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func makeHelpCallback(reg map[string]cliCommand) func(cfg *config.Config) error {
	return func(cfg *config.Config) error {
		fmt.Println("Welcome to the Pokedex!")
        fmt.Println("Usage:\n")
        // iterate the captured reg
        for _, cmd := range reg {
            fmt.Printf("%s: %s\n", cmd.name, cmd.description)
        }
        return nil
	}
}

