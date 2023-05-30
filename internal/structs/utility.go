package structs

type Language struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Official bool   `json:"official"`
	ISO639   string `json:"iso639"`
	ISO3166  string `json:"iso3166"`
	Names    []Name `json:"names"`
}
type APIResource struct {
	URL string `json:"url"`
}

type Description struct {
	Description string           `json:"description"`
	Language    NamedAPIResource `json:"language"`
}

type Effect struct {
	Effect   string           `json:"description"`
	Language NamedAPIResource `json:"language"`
}

type Encounter struct {
	MinLevel        int                `json:"min_level"`
	MaxLevel        int                `json:"max_level"`
	ConditionValues []NamedAPIResource `json:"condition_values"`
	Chance          int                `json:"chance"`
	Method          []NamedAPIResource `json:"method"`
}

type FlavorText struct {
	FlavorText string             `json:"flavor_text"`
	Language   []NamedAPIResource `json:"language"`
	Version    []NamedAPIResource `json:"version"`
}

type GenerationGameIndex struct {
	GameIndex  int              `json:"game_index"`
	Generation NamedAPIResource `json:"generation"`
}

type MachineVersionDetail struct {
	Machine      APIResource      `json:"machine"`
	VersionGroup NamedAPIResource `json:"version_group"`
}

type Name struct {
	Name     string `json:"name"`
	Language string `json:"version_group"`
}

type NamedAPIResource struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type VerboseEffect struct {
	Effect      string           `json:"effect"`
	ShortEffect string           `json:"short_effect"`
	Language    NamedAPIResource `json:"language"`
}

type VersionEncounterDetail struct {
	Version          NamedAPIResource `json:"version"`
	MaxChance        int              `json:"max_chance"`
	EncounterDetails []Encounter      `json:"encounter_details"`
}

type VersionGroupFlavorText struct {
	Text         string           `json:"text"`
	Language     NamedAPIResource `json:"language"`
	VersionGroup NamedAPIResource `json:"version_group"`
}
