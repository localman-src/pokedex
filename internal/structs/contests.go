package structs

type ContestType struct {
	ID          int              `json:"id"`
	Name        string           `json:"Name"`
	BerryFlavor NamedAPIResource `json:"berry_flavor"`
	Names       []ContestName    `json:"names"`
}

type ContestName struct {
	Name     string           `json:"id"`
	Color    string           `json:"color"`
	Language NamedAPIResource `json:"language"`
}

type ContestEffect struct {
	ID                int          `json:"id"`
	Appeal            int          `json:"appeal"`
	Jam               int          `json:"jam"`
	EffectEntries     []Effect     `json:"effect_entries"`
	FlavorTextEntries []FlavorText `json:"flavor_text_entries"`
}

type SuperContestEffect struct {
	ID                int                `json:"id"`
	Appeal            int                `json:"appeal"`
	FlavorTextEntries []FlavorText       `json:"flavor_text_entries"`
	Moves             []NamedAPIResource `json:"moves"`
}
