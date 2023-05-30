package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type locationArea struct {
	ID                   int    `json:"id"`
	Name                 string `json:"name"`
	GameIndex            int    `json:"game_index"`
	EncounterMethodRates []struct {
		EncounterMethod map[string]string `json:"encounter_method"`
		VersionDetails  []struct {
			Rate    int                 `json:"rate"`
			Version []map[string]string `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	Location          map[string]string `json:"location"`
	Names             interface{}       `json:"names"`
	PokemonEncounters interface{}       `json:"pokemon_encounters"`
}

type locationAreas []locationArea

type locationAreaResponse struct {
	Count    int                 `json:"count"`
	Next     string              `json:"next"`
	Previous string              `json:"previous"`
	Results  []map[string]string `json:"results"`
}

func (l locationAreaResponse) PrintAreas() {
	for _, location := range l.Results {
		println(location["name"])
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
