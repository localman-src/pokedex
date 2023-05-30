package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/localman-src/pokedex/internal/structs"
)

type locationAreaResponse struct {
	Count    int                        `json:"count"`
	Next     string                     `json:"next"`
	Previous string                     `json:"previous"`
	Results  []structs.NamedAPIResource `json:"results"`
}

func (l locationAreaResponse) PrintAreas() {
	for _, location := range l.Results {
		println(location.Name)
	}
}
func GetAreas(url string) (*locationAreaResponse, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error fetching areas: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}
	locationResponse := locationAreaResponse{}

	err = json.Unmarshal(body, &locationResponse)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling JSML: %v", err)
	}

	return &locationResponse, nil
}
