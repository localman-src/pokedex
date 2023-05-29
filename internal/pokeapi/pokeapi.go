package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type locationArea struct {
	ID                   int               `json:"id"`
	Name                 string            `json:"name"`
	GameIndex            int               `json:"game_index"`
	EncounterMethodRates interface{}       `json:"encounter_method_rates"`
	Location             map[string]string `json:"location"`
	Names                interface{}       `json:"names"`
	PokemonEncounters    interface{}       `json:"pokemon_encounters"`
}

type locationAreas []locationArea

type locationAreaResponse struct {
	Count    int                 `json:"count"`
	Next     string              `json:"next"`
	Previous string              `json:"previous"`
	Results  []map[string]string `json:"results"`
}

func GetAreas() (locationAreaResponse, error) {
	resp, err := http.Get("https://pokeapi.co/api/v2/location-area")
	if err != nil {
		return nil, fmt.Errorf("error fetching areas: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}
	locations := locationAreaResponse{}

	err = json.Unmarshal(body, &locations)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling JSML: %v", err)
	}

	return locations, nil
}
