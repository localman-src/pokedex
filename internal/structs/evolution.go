package structs

type EvolutionChain struct {
	ID              int                      `json:"id"`
	BabyTriggerItem []NamedAPIResource[Item] `json:"baby_trigger_item"`
	Chain           ChainLink                `json:"chain"`
}

type ChainLink struct {
	IsBaby           bool                             `json:"is_baby"`
	Species          NamedAPIResource[PokemonSpecies] `json:"species"`
	EvolutionDetails []EvolutionDetail                `json:"evolution_details"`
	EvolvesTo        []ChainLink                      `json:"evolves_to"`
}

type EvolutionDetail struct {
	Item                  NamedAPIResource[Item]             `json:"item"`
	Trigger               NamedAPIResource[EvolutionTrigger] `json:"trigger"`
	Gender                int                                `json:"gender"`
	HeldItem              NamedAPIResource[Item]             `json:"held_item"`
	KnownMove             NamedAPIResource[Move]             `json:"known_move"`
	KnownMoveType         NamedAPIResource[Type]             `json:"known_move_type"`
	Location              NamedAPIResource[Location]         `json:"location"`
	MinLevel              int                                `json:"min_level"`
	MinHappiness          int                                `json:"min_happiness"`
	MinBeauty             int                                `json:"min_beauty"`
	MinAffection          int                                `json:"min_affection"`
	NeedsOverworldRain    bool                               `json:"needs_overworld_rain"`
	PartySpecies          NamedAPIResource[PokemonSpecies]   `json:"party_species"`
	PartyType             NamedAPIResource[Type]             `json:"party_type"`
	RelativePhysicalStats int                                `json:"relative_physical_stats"`
	TimeOfDay             string                             `json:"time_of_day"`
	TradeSpecies          NamedAPIResource[PokemonSpecies]   `json:"trade_species"`
	TurnUpsideDown        bool                               `json:"turn_upside_down"`
}

type EvolutionTrigger struct {
	ID             int                              `json:"id"`
	Name           string                           `json:"name"`
	Names          []Name                           `json:"names"`
	PokemonSpecies NamedAPIResource[PokemonSpecies] `json:"pokemon_species"`
}
