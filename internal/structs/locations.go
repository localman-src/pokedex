package structs

type Location struct {
	ID          string                `json:"id"`
	Name        string                `json:"name"`
	Region      NamedAPIResource      `json:"region"`
	Names       []Name                `json:"names"`
	GameIndices []GenerationGameIndex `json:"game_indices"`
	Areas       []NamedAPIResource    `json:"areas"`
}

type LocationArea struct {
	ID                   int                   `json:"ID"`
	Name                 string                `json:"name"`
	GameIndex            int                   `json:"game_index"`
	EncounterMethodRates []EncounterMethodRate `json:"encounter_method_rates"`
	Location             NamedAPIResource      `json:"location"`
	Names                []Name                `json:"names"`
	PokemonEncounters    []PokemonEncounter    `json:"pokemon_encounters"`
}

type EncounterMethodRate struct {
	EncounterMethod NamedAPIResource          `json:"encounter_method"`
	VersionDetails  []EncounterVersionDetails `json:"version_details"`
}

type EncounterVersionDetails struct {
	Rate    int              `json:"rate"`
	Version NamedAPIResource `json:"version"`
}

type PokemonEncounter struct {
	Name           NamedAPIResource         `json:"name"`
	VersionDetails []VersionEncounterDetail `json:"version_details"`
}

type PalParkArea struct {
	ID                int                       `json:"id"`
	Name              string                    `json:"name"`
	Names             []Name                    `json:"names"`
	PokemonEncounters []PalParkEncounterSpecies `json:"pokemon_encounters"`
}

type PalParkEncounterSpecies struct {
	BaseScore      int              `json:"id"`
	Rate           int              `json:"rate"`
	PokemonSpecies NamedAPIResource `json:"pokemon_species"`
}

type Region struct {
	ID             int                `json:"id"`
	Locations      []NamedAPIResource `json:"locations"`
	Name           string             `json:"name"`
	Names          []Name             `json:"names"`
	MainGeneration []NamedAPIResource `json:"main_generation"`
	Pokedexes      []NamedAPIResource `json:"pokedexes"`
	VersionGroups  []NamedAPIResource `json:"version_groups"`
}
