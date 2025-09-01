package main

import (
	"pokedexcli/internal/config"
	"pokedexcli/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config.Config) error
}

func main() {
	reg := map[string]cliCommand{} // step 1: create

	reg["exit"] = cliCommand{
	    name:        "exit",
	    description: "Exit the Pokedex",
	    callback:    commandExit,
	}

	reg["help"] = cliCommand{
	    name:        "help",
	    description: "Displays a help message",
	    callback:    makeHelpCallback(reg), // step 2: now reg exists
	}

	reg["map"] = cliCommand{
		name:        "map",
		description: "Displays the names of the next 20 location areas in the Pokemon world",
		callback:    pokeapi.Map,
	}

	reg["mapb"] = cliCommand{
		name:        "mapb",
		description: "Displays the names of the previous 20 location areas in the Pokemon world",
		callback:    pokeapi.Mapb,
	}
	run(reg)
}