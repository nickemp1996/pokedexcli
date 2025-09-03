package pokeapi

import (
	"fmt"
	"io"
	"time"
	"net/http"
	"crypto/tls"
	"pokedexcli/internal/pokecache"
)

var C *pokecache.Cache = pokecache.NewCache(30 * time.Second)

func request(url string) ([]byte, error) {
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
            return nil, err
        }
        defer res.Body.Close()

        if res.StatusCode > 299 {
            b, _ := io.ReadAll(res.Body)
            return nil, fmt.Errorf("response failed with status code: %d and body: %s", res.StatusCode, string(b))
        }

        body, err = io.ReadAll(res.Body) // assign, not :=
        if err != nil {
            return nil, err
        }

        C.Add(url, body)
    }

    return body, nil
}