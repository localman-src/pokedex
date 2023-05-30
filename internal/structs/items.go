package structs

type Item struct {
	ID                int                      `json:"id"`
	Name              string                   `json:"name"`
	Cost              int                      `json:"cost"`
	FlingPower        int                      `json:"fling_power"`
	FlingEffect       NamedAPIResource         `json:"fling_effect"`
	Attributes        []NamedAPIResource       `json:"attributes"`
	Category          NamedAPIResource         `json:"category"`
	EffectEntries     []VerboseEffect          `json:"effect_entries"`
	FlavorTextEntries []VersionGroupFlavorText `json:"flavor_text_entries"`
	GameIndices       []GenerationGameIndex    `json:"game_indices"`
	Names             []Name                   `json:"names"`
	Sprites           ItemSprites              `json:"sprites"`
	HeldByPokemon     []ItemHolderPokemon      `json:"held_by_pokemon"`
	BabyTriggerFor    APIResource              `json:"baby_trigger_for"`
	Machines          []MachineVersionDetail   `json:"machines"`
}

type ItemSprites struct {
	Default string `json:"default"`
}

type ItemHolderPokemon struct {
	Pokemon        NamedAPIResource                 `json:"pokemon"`
	VersionDetails []ItemHolderPokemonVersionDetail `json:"version_details"`
}

type ItemHolderPokemonVersionDetail struct {
	Rarity  int              `json:"rarity"`
	Version NamedAPIResource `json:"version"`
}

type ItemAttribute struct {
	ID           int                `json:"id"`
	Name         string             `json:"name"`
	Items        []NamedAPIResource `json:"items"`
	Names        []Name             `json:"names"`
	Descriptions []Description      `json:"descriptions"`
}

type ItemCategory struct {
	ID     int                `json:"id"`
	Name   string             `json:"name"`
	Items  []NamedAPIResource `json:"items"`
	Names  []Name             `json:"names"`
	Pocket NamedAPIResource   `json:"pocket"`
}

type ItemFlingEffect struct {
	ID            int                `json:"id"`
	Name          string             `json:"name"`
	EffectEntries []Effect           `json:"effect_entries"`
	Items         []NamedAPIResource `json:"items"`
}

type ItemPocket struct {
	ID         int                `json:"id"`
	Name       string             `json:"name"`
	Categories []NamedAPIResource `json:"categories"`
	Names      []Name             `json:"names"`
}
