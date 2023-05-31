package structs

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Language struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Official bool   `json:"official"`
	ISO639   string `json:"iso639"`
	ISO3166  string `json:"iso3166"`
	Names    []Name `json:"names"`
}
type APIResource[T any] struct {
	URL string `json:"url"`
}

func (r *APIResource[T]) GetPokeAPI() (resourceObject *T, err error) {
	response, err := http.Get(r.URL)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	err = json.Unmarshal(body, &resourceObject)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling JSON: %v", err)
	}

	return resourceObject, err

}

type Description struct {
	Description string                     `json:"description"`
	Language    NamedAPIResource[Language] `json:"language"`
}

type Effect struct {
	Effect   string                     `json:"description"`
	Language NamedAPIResource[Language] `json:"language"`
}

type Encounter struct {
	MinLevel        int                                         `json:"min_level"`
	MaxLevel        int                                         `json:"max_level"`
	ConditionValues []NamedAPIResource[EncounterConditionValue] `json:"condition_values"`
	Chance          int                                         `json:"chance"`
	Method          NamedAPIResource[EncounterMethod]           `json:"method"`
}

type FlavorText struct {
	FlavorText string                       `json:"flavor_text"`
	Language   []NamedAPIResource[Language] `json:"language"`
	Version    []NamedAPIResource[Version]  `json:"version"`
}

type GenerationGameIndex struct {
	GameIndex  int                          `json:"game_index"`
	Generation NamedAPIResource[Generation] `json:"generation"`
}

type MachineVersionDetail struct {
	Machine      APIResource[Machine]           `json:"machine"`
	VersionGroup NamedAPIResource[VersionGroup] `json:"version_group"`
}

type Name struct {
	Name     string                     `json:"name"`
	Language NamedAPIResource[Language] `json:"language"`
}

type NamedAPIResource[T any] struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func (r *NamedAPIResource[T]) Get() (response T) {

	return response
}

type NamedAPIResourceList[T any] struct {
	Count    int                   `json:"count"`
	Next     string                `json:"next"`
	Previous string                `json:"previous"`
	Results  []NamedAPIResource[T] `json:"results"`
}

func (r *NamedAPIResourceList[T]) Print() {
	for _, resource := range r.Results {
		println(resource.Name)
	}
}

type VerboseEffect struct {
	Effect      string                     `json:"effect"`
	ShortEffect string                     `json:"short_effect"`
	Language    NamedAPIResource[Language] `json:"language"`
}

type VersionEncounterDetail struct {
	Version          NamedAPIResource[Version] `json:"version"`
	MaxChance        int                       `json:"max_chance"`
	EncounterDetails []Encounter               `json:"encounter_details"`
}

type VersionGameIndex struct {
	GameIndex int                       `json:"game_index"`
	Version   NamedAPIResource[Version] `json:"version"`
}
type VersionGroupFlavorText struct {
	Text         string                         `json:"text"`
	Language     NamedAPIResource[Language]     `json:"language"`
	VersionGroup NamedAPIResource[VersionGroup] `json:"version_group"`
}
