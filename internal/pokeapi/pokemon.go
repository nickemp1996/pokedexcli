package pokeapi

import (
	"encoding/json"
	"pokedexcli/internal/structs"
)

func GetPokemon(name string) (*structs.Pokemon, error) {
	url := "https://pokeapi.co/api/v2/pokemon/" + name

	body, err := request(url)
	if err != nil {
		return nil, err
	}

	var pokemon structs.Pokemon
    if err := json.Unmarshal(body, &pokemon); err != nil {
        return nil, err
    }

    return &pokemon, nil
}