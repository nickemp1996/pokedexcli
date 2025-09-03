package main

import (
	"pokedexcli/internal/config"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config.Config, []string) error
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
		callback:    commandMap,
	}

	reg["mapb"] = cliCommand{
		name:        "mapb",
		description: "Displays the names of the previous 20 location areas in the Pokemon world",
		callback:    commandMapb,
	}

	reg["explore"] = cliCommand{
		name:        "explore",
		description: "Displays the names of the Pokemon in a given location area",
		callback:    commandExplore,
	}

	reg["catch"] = cliCommand{
		name:        "catch",
		description: "Gives the user the opportunity to try and catch the named Pokemon",
		callback:    commandCatch,
	}

	reg["inspect"] = cliCommand{
		name:        "inspect",
		description: "Displays details about the named Pokemon if the user has already caught that Pokemon",
		callback:    commandInspect,
	}

	reg["pokedex"] = cliCommand{
		name:        "pokedex",
		description: "Lists the names of all the Pokemon the user has caught",
		callback:    commandPokedex,
	}
	run(reg)
}