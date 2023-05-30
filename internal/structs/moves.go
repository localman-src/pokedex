package structs

type Move struct {
	ID                 int                    `json:"id"`
	Name               string                 `json:"name"`
	Accuracy           int                    `json:"accuracy"`
	EffectChance       int                    `json:"effect_chance"`
	PP                 int                    `json:"pp"`
	Priority           int                    `json:"priority"`
	Power              int                    `json:"power"`
	ContestCombos      ContestComboSets       `json:"contest_combos"`
	ContestType        NamedAPIResource       `json:"contest_type"`
	ContestEffect      APIResource            `json:"contest_effect"`
	DamageClass        NamedAPIResource       `json:"damage_class"`
	EffectEntries      []VerboseEffect        `json:"effect_entries"`
	EfectChanges       []AbilityEffectChange  `json:"effect_changes"`
	LearnedByPokemon   []NamedAPIResource     `json:"learned_by_pokemon"`
	FlavorTextEntries  []MoveFlavorText       `json:"flavor_text_entries"`
	Generation         NamedAPIResource       `json:"generation"`
	Machines           []MachineVersionDetail `json:"machines"`
	Meta               MoveMetaData           `json:"meta"`
	Names              []Name                 `json:"names"`
	PastValues         []PastMoveStatValues   `json:"past_values"`
	StatChanges        []MoveStatChange       `json:"stat_changes"`
	SuperContestEffect APIResource            `json:"super_contest_effect"`
	Target             NamedAPIResource       `json:"target"`
	Type               NamedAPIResource       `json:"type"`
}

type ContestComboSets struct {
	Normal ContestComboDetail `json:"normal"`
	Super  ContestComboDetail `json:"super"`
}

type ContestComboDetail struct {
	UseBefore []NamedAPIResource `json:"use_before"`
	UseAfter  []NamedAPIResource `json:"id"`
}

type MoveFlavorText struct {
	FlavorText   string           `json:"flavor_text"`
	Language     NamedAPIResource `json:"language"`
	VersionGroup NamedAPIResource `json:"version_group"`
}

type MoveMetaData struct {
}
