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
	"pokedexcli/internal/pokecache"
)

var C *pokecache.Cache = pokecache.NewCache(30 * time.Second)

func getLocationAreas(url string, cfg *config.Config) error {
    body, ok := C.Get(url)
    if !ok {
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
        defer res.Body.Close()

        if res.StatusCode > 299 {
            b, _ := io.ReadAll(res.Body)
            return fmt.Errorf("response failed with status code: %d and body: %s", res.StatusCode, string(b))
        }

        body, err = io.ReadAll(res.Body) // assign, not :=
        if err != nil {
            return err
        }

        C.Add(url, body)
    }

    var locationAreaList structs.LocationAreaList
    if err := json.Unmarshal(body, &locationAreaList); err != nil {
        return err
    }

    for _, r := range locationAreaList.Results {
        fmt.Println(r.Name)
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