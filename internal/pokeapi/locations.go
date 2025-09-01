package pokeapi

import (
	"fmt"
	"io"
	"time"
	"crypto/tls"
	"encoding/json"
	"net/http"
	"pokedexcli/internal/structs"
	"pokedexcli/internal/config"
)

func getLocationAreas(url string, cfg *config.Config) error {
	client := &http.Client{
		Timeout: 15 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{MinVersion: tls.VersionTLS12},
		},
	}
	res, err := client.Get(url)
	if err != nil {
		return err
	}

	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if res.StatusCode > 299 {
		return fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		return err
	}

	locationAreaList := structs.LocationAreaList{}
	err = json.Unmarshal(body, &locationAreaList)
	if err != nil {
		return err
	}

	for _, result := range locationAreaList.Results {
		fmt.Println(result.Name)
	}

	cfg.Next = locationAreaList.Next
	cfg.Previous = locationAreaList.Previous

	return nil
}

func Map(cfg *config.Config) error {
	url := ""
	if cfg.Next != nil {
		url = *cfg.Next
	} else {
		url = "https://pokeapi.co/api/v2/location-area"
	}
	
	err := getLocationAreas(url, cfg)
	if err != nil {
		return err
	}

	return nil
}

func Mapb(cfg *config.Config) error {
	url := ""
	if cfg.Previous != nil {
		url = *cfg.Previous
	} else {
		fmt.Println("You are on the first page")
		return nil
	}

	err := getLocationAreas(url, cfg)
	if err != nil {
		return err
	}

	return nil
}