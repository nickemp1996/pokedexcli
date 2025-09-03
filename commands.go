package main

import (
	"fmt"
	"os"
	"math/rand"
	"pokedexcli/internal/config"
	"pokedexcli/internal/pokeapi"
	"pokedexcli/internal/structs"
)

var Pokedex = map[string]structs.Pokemon{}

func commandExit(cfg *config.Config, params []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func makeHelpCallback(reg map[string]cliCommand) func(cfg *config.Config, params []string) error {
	return func(cfg *config.Config, params []string) error {
		fmt.Println("Welcome to the Pokedex!")
        fmt.Println("Usage:")
        // iterate the captured reg
        for _, cmd := range reg {
            fmt.Printf("%s: %s\n", cmd.name, cmd.description)
        }
        return nil
	}
}

func commandMap(cfg *config.Config, params []string) error {
	url := ""
	if cfg.Next != nil {
		url = *cfg.Next
	} else {
		url = "https://pokeapi.co/api/v2/location-area"
	}

	locationAreaList, err := pokeapi.GetLocationAreas(url)
	if err != nil {
		return err
	}

	for _, r := range locationAreaList.Results {
        fmt.Println(r.Name)
    }

    cfg.Next = locationAreaList.Next
    cfg.Previous = locationAreaList.Previous

    return nil
}

func commandMapb(cfg *config.Config, params []string) error {
	url := ""
	if cfg.Previous != nil {
		url = *cfg.Previous
	} else {
		fmt.Println("You are on the first page")
		return nil
	}

	locationAreaList, err := pokeapi.GetLocationAreas(url)
	if err != nil {
		return err
	}

	for _, r := range locationAreaList.Results {
        fmt.Println(r.Name)
    }

    cfg.Next = locationAreaList.Next
    cfg.Previous = locationAreaList.Previous

    return nil
}

func commandExplore(cfg *config.Config, params []string) error {
	if len(params) == 0 {
		return fmt.Errorf("Location area to explore was not provided...")
	}

	fmt.Printf("Exploring %s...\n", params[0])

	locationArea, err := pokeapi.GetLocationArea(params[0])
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, pokemonEncounter := range locationArea.PokemonEncounters {
		fmt.Println(" -", pokemonEncounter.Pokemon.Name)
	}

	return nil
}

func commandCatch(cfg *config.Config, params []string) error {
	if len(params) == 0 {
		return fmt.Errorf("Pokemon to catch was not provided...")
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", params[0])

	r := rand.Float64()
	pokemon, err := pokeapi.GetPokemon(params[0])
	if err != nil {
		return err
	}

	p := 20.0/float64(pokemon.BaseExperience)
	if r < p {
		fmt.Println(pokemon.Name, "was caught!")
		Pokedex[pokemon.Name] = *pokemon
	} else {
		fmt.Println(pokemon.Name, "escaped!")
	}

	return nil
}

func commandInspect(cfg *config.Config, params []string) error {
	if len(params) == 0 {
		return fmt.Errorf("Pokemon to inspect was not provided...")
	}

	pokemon, ok := Pokedex[params[0]]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	fmt.Println("Height:", pokemon.Height)
	fmt.Println("Weight:", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  - %s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, poketype := range pokemon.Types {
		fmt.Println("  -", poketype.Type.Name)
	}

	return nil
}