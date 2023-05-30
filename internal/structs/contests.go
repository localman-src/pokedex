package structs

type ContestType struct {
	ID          int                           `json:"id"`
	Name        string                        `json:"Name"`
	BerryFlavor NamedAPIResource[BerryFlavor] `json:"berry_flavor"`
	Names       []ContestName                 `json:"names"`
}

type ContestName struct {
	Name     string                     `json:"id"`
	Color    string                     `json:"color"`
	Language NamedAPIResource[Language] `json:"language"`
}

type ContestEffect struct {
	ID                int          `json:"id"`
	Appeal            int          `json:"appeal"`
	Jam               int          `json:"jam"`
	EffectEntries     []Effect     `json:"effect_entries"`
	FlavorTextEntries []FlavorText `json:"flavor_text_entries"`
}

type SuperContestEffect struct {
	ID                int                      `json:"id"`
	Appeal            int                      `json:"appeal"`
	FlavorTextEntries []FlavorText             `json:"flavor_text_entries"`
	Moves             []NamedAPIResource[Move] `json:"moves"`
}
