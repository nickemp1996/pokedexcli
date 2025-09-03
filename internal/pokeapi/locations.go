package pokeapi

import (
	"encoding/json"
	"pokedexcli/internal/structs"
)

func GetLocationArea(areaName string) (*structs.LocationArea, error) {
	url := "https://pokeapi.co/api/v2/location-area/" + areaName

	body, err := request(url)
	if err != nil {
		return nil, err
	}

    var locationArea structs.LocationArea
    if err := json.Unmarshal(body, &locationArea); err != nil {
        return nil, err
    }

    return &locationArea, nil
}

func GetLocationAreas(url string) (*structs.LocationAreaList, error) {
	body, err := request(url)
	if err != nil {
		return nil, err
	}

    var locationAreaList structs.LocationAreaList
    if err := json.Unmarshal(body, &locationAreaList); err != nil {
        return nil, err
    }

    return &locationAreaList, nil
}