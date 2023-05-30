package structs

type Generation struct {
	ID             int                                `json:"id"`
	Name           string                             `json:"name"`
	Abilities      []NamedAPIResource[Ability]        `json:"abilities"`
	Names          []Name                             `json:"names"`
	MainRegion     []NamedAPIResource[Region]         `json:"main_region"`
	Moves          []NamedAPIResource[Move]           `json:"moves"`
	PokemonSpecies []NamedAPIResource[PokemonSpecies] `json:"pokemon_species"`
	Types          []NamedAPIResource[Type]           `json:"types"`
	VersionGroups  []NamedAPIResource[VersionGroup]   `json:"version_groups"`
}

type Pokedex struct {
	ID             int                              `json:"id"`
	Name           string                           `json:"name"`
	IsMainSeries   bool                             `json:"is_main_series"`
	Descriptions   []Description                    `json:"descriptions"`
	Names          []Name                           `json:"names"`
	PokemonEntries []PokemonEntry                   `json:"pokemon_entries"`
	Region         NamedAPIResource[Region]         `json:"region"`
	VersionGroups  []NamedAPIResource[VersionGroup] `json:"version_groups"`
}

type PokemonEntry struct {
	EntryNumber    int                              `json:"entry_number"`
	PokemonSpecies NamedAPIResource[PokemonSpecies] `json:"pokemon_species"`
}

type Version struct {
	ID           int                            `json:"id"`
	Name         string                         `json:"name"`
	Names        []Name                         `json:"names"`
	VersionGroup NamedAPIResource[VersionGroup] `json:"version_group"`
}

type VersionGroup struct {
	ID               int                                 `json:"id"`
	Name             string                              `json:"name"`
	Order            int                                 `json:"order"`
	Generation       NamedAPIResource[Generation]        `json:"generation"`
	MoveLearnMethods []NamedAPIResource[MoveLearnMethod] `json:"move_learn_methods"`
	Pokedexes        []NamedAPIResource[Pokedex]         `json:"pokedexes"`
	Regions          []NamedAPIResource[Region]          `json:"regions"`
	Versions         []NamedAPIResource[Version]         `json:"versions"`
}
