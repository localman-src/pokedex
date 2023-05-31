package structs

import "fmt"

type Location struct {
	ID          string                           `json:"id"`
	Name        string                           `json:"name"`
	Region      NamedAPIResource[Region]         `json:"region"`
	Names       []Name                           `json:"names"`
	GameIndices []GenerationGameIndex            `json:"game_indices"`
	Areas       []NamedAPIResource[LocationArea] `json:"areas"`
}

type LocationArea struct {
	ID                   int                        `json:"id"`
	Name                 string                     `json:"name"`
	GameIndex            int                        `json:"game_index"`
	EncounterMethodRates []EncounterMethodRate      `json:"encounter_method_rates"`
	Location             NamedAPIResource[Location] `json:"location"`
	Names                []Name                     `json:"names"`
	PokemonEncounters    []PokemonEncounter         `json:"pokemon_encounters"`
}

func (l *LocationArea) PrintPokemonEncounters() {
	fmt.Printf("Found these pokemon in %s:\n", l.Name)
	for _, encounter := range l.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}
}

type EncounterMethodRate struct {
	EncounterMethod NamedAPIResource[EncounterMethod] `json:"encounter_method"`
	VersionDetails  []EncounterVersionDetails         `json:"version_details"`
}

type EncounterVersionDetails struct {
	Rate    int                       `json:"rate"`
	Version NamedAPIResource[Version] `json:"version"`
}

type PokemonEncounter struct {
	Pokemon        NamedAPIResource[Pokemon] `json:"pokemon"`
	VersionDetails []VersionEncounterDetail  `json:"version_details"`
}

type PalParkArea struct {
	ID                int                       `json:"id"`
	Name              string                    `json:"name"`
	Names             []Name                    `json:"names"`
	PokemonEncounters []PalParkEncounterSpecies `json:"pokemon_encounters"`
}

type PalParkEncounterSpecies struct {
	BaseScore      int                              `json:"id"`
	Rate           int                              `json:"rate"`
	PokemonSpecies NamedAPIResource[PokemonSpecies] `json:"pokemon_species"`
}

type Region struct {
	ID             int                              `json:"id"`
	Locations      []NamedAPIResource[Location]     `json:"locations"`
	Name           string                           `json:"name"`
	Names          []Name                           `json:"names"`
	MainGeneration []NamedAPIResource[Generation]   `json:"main_generation"`
	Pokedexes      []NamedAPIResource[Pokedex]      `json:"pokedexes"`
	VersionGroups  []NamedAPIResource[VersionGroup] `json:"version_groups"`
}
