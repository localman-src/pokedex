package structs

type Ability struct {
	ID                int                          `json:"id"`
	Name              string                       `json:"name"`
	IsMainSeries      bool                         `json:"is_main_series"`
	Generation        NamedAPIResource[Generation] `json:"generation"`
	Names             []Name                       `json:"names"`
	EffectEntries     []VerboseEffect              `json:"effect_entries"`
	EffectChanges     []AbilityEffectChange        `json:"effect_changes"`
	FlavorTextEntries []AbilityFlavorText          `json:"flavor_text_entries"`
	Pokemon           []AbilityPokemon             `json:"pokemon"`
}

type AbilityEffectChange struct {
	EffectEntries []Effect                       `json:"effect_entries"`
	VersionGroup  NamedAPIResource[VersionGroup] `json:"version_group"`
}

type AbilityFlavorText struct {
	FlavorText   string                         `json:"flavor_text"`
	Language     NamedAPIResource[Language]     `json:"language"`
	VersionGroup NamedAPIResource[VersionGroup] `json:"version_group"`
}

type AbilityPokemon struct {
	IsHidden bool                      `json:"is_hidden"`
	Slot     int                       `json:"slot"`
	Pokemon  NamedAPIResource[Pokemon] `json:"pokemon"`
}

type Characteristic struct {
	ID             int   `json:"id"`
	GeneModulo     int   `json:"gene_modulo"`
	PossibleValues []int `json:"possible_values"`
}

type EggGroup struct {
	ID             int                                `json:"id"`
	Name           string                             `json:"name"`
	Names          []Name                             `json:"names"`
	PokemonSpecies []NamedAPIResource[PokemonSpecies] `json:"pokemon_species"`
}

type Gender struct {
	ID                    int                                `json:"id"`
	Name                  string                             `json:"name"`
	PokemonSpeciesDetails []PokemonSpeciesGender             `json:"pokemon_species_details"`
	RequiredForEvolution  []NamedAPIResource[PokemonSpecies] `json:"required_for_evolution"`
}

type PokemonSpeciesGender struct {
	Rate           int                              `json:"rate"`
	PokemonSpecies NamedAPIResource[PokemonSpecies] `json:"pokemon_species"`
}

type GrowthRate struct {
	ID             int                                `json:"id"`
	Name           string                             `json:"name"`
	Formula        string                             `json:"formula"`
	Descriptions   []Description                      `json:"descriptions"`
	Levels         []GrowthRateExperienceLevel        `json:"levels"`
	PokemonSpecies []NamedAPIResource[PokemonSpecies] `json:"pokemon_species"`
}

type GrowthRateExperienceLevel struct {
	Level      int `json:"level"`
	Experience int `json:"experience"`
}

type Nature struct {
	ID                         int                           `json:"id"`
	Name                       string                        `json:"name"`
	DecreasedState             NamedAPIResource[Stat]        `json:"decreased_stat"`
	IncreasedStat              NamedAPIResource[Stat]        `json:"increased_stat"`
	HatesFlavor                NamedAPIResource[BerryFlavor] `json:"hates_flavor"`
	LikesFlavor                NamedAPIResource[BerryFlavor] `json:"likes_flavor"`
	PokeathlonStatChanges      []NatureStatChange            `json:"pokeathlon_stat_changes"`
	MoveBattleStylePreferences []MoveBattleStylePreference   `json:"move_battle_style_preferences"`
	Names                      []Name                        `json:"names"`
}

type NatureStatChange struct {
	MaxChange      int                              `json:"max_change"`
	PokeathlonStat NamedAPIResource[PokeathlonStat] `json:"pokeathlon_stat"`
}

type MoveBattleStylePreference struct {
	LowHpPreference  int                               `json:"low_hp_preference"`
	HighHpPreference int                               `json:"high_hp_preference"`
	MoveBattleStyle  NamedAPIResource[MoveBattleStyle] `json:"move_battle_style"`
}

type PokeathlonStat struct {
	ID               int                            `json:"id"`
	Name             string                         `json:"name"`
	Named            []Name                         `json:"names"`
	AffectingNatures NaturePokeathlonStatAffectSets `json:"affecting_natures"`
}

type NaturePokeathlonStatAffectSets struct {
	Increase []NaturePokeathlonStatAffect `json:"increase"`
	Decrease []NaturePokeathlonStatAffect `json:"decrease"`
}

type NaturePokeathlonStatAffect struct {
	MaxChange int                      `json:"max_change"`
	Nature    NamedAPIResource[Nature] `json:"nature"`
}

type Pokemon struct {
	ID                     int                              `json:"id"`
	Name                   string                           `json:"name"`
	BaseExperience         int                              `json:"base_experience"`
	Height                 int                              `json:"height"`
	IsDefault              bool                             `json:"is_default"`
	Order                  int                              `json:"order"`
	Weight                 int                              `json:"weight"`
	Abilities              []PokemonAbility                 `json:"abilities"`
	Forms                  []NamedAPIResource[PokemonForm]  `json:"forms"`
	GameIndices            []VersionGameIndex               `json:"game_indices"`
	HeldItems              []PokemonHeldItem                `json:"held_items"`
	LocationAreaEncounters string                           `json:"location_area_encounters"`
	Moves                  []PokemonMove                    `json:"moves"`
	PastTypes              []PokemonTypePast                `json:"past_types"`
	Sprites                PokemonSprites                   `json:"sprites"`
	Species                NamedAPIResource[PokemonSpecies] `json:"species"`
	Stats                  []PokemonStat                    `json:"stats"`
	Types                  []PokemonType                    `json:"types"`
}

type PokemonAbility struct {
	IsHidden bool                      `json:"is_hidden"`
	Slot     int                       `json:"slot"`
	Ability  NamedAPIResource[Ability] `json:"ability"`
}

type PokemonType struct {
	Slot int                    `json:"slot"`
	Type NamedAPIResource[Type] `json:"type"`
}

type PokemonFormType struct {
	Slot int                    `json:"slot"`
	Type NamedAPIResource[Type] `json:"type"`
}

type PokemonTypePast struct {
	Generation NamedAPIResource[Generation] `json:"generation"`
	Types      []PokemonType                `json:"types"`
}

type PokemonHeldItem struct {
	Item           NamedAPIResource[Item]   `json:"item"`
	VersionDetails []PokemonHeldItemVersion `json:"version_details"`
}

type PokemonHeldItemVersion struct {
	Version NamedAPIResource[Version] `json:"version"`
	Rarity  int                       `json:"rarity"`
}

type PokemonMove struct {
	Move                NamedAPIResource[Move] `json:"move"`
	VersionGroupDetails []PokemonMoveVersion   `json:"version_group_details"`
}

type PokemonMoveVersion struct {
	MoveLearnMethod NamedAPIResource[MoveLearnMethod] `json:"move_learn_method"`
	VersionGroup    NamedAPIResource[VersionGroup]    `json:"version_group"`
	LevelLearnedAt  int                               `json:"level_learned_at"`
}

type PokemonStat struct {
	Stat     NamedAPIResource[Stat] `json:"stat"`
	Effort   int                    `json:"effort"`
	BaseStat int                    `json:"base_stat"`
}

type PokemonSprites struct {
	FrontDefault     string `json:"front_default"`
	FrontShiny       string `json:"front_shiny"`
	FrontFemale      string `json:"front_female"`
	FrontShinyFemale string `json:"front_shiny_female"`
	BackDefault      string `json:"back_default"`
	BackShiny        string `json:"back_shiny"`
	BackFemale       string `json:"back_female"`
	BackShinyFemale  string `json:"back_shiny_female"`
}

type LocationAreaEncounter struct {
	LocationArea   NamedAPIResource[LocationArea] `json:"location_area"`
	VersionDetails []VersionEncounterDetail       `json:"version_details"`
}

type PokemonColor struct {
	ID             int                                `json:"id"`
	Name           string                             `json:"name"`
	Names          []Name                             `json:"names"`
	PokemonSpecies []NamedAPIResource[PokemonSpecies] `json:"pokemon_species"`
}

type PokemonForm struct {
	ID           int                            `json:"id"`
	Name         string                         `json:"name"`
	Order        int                            `json:"order"`
	FormOrder    int                            `json:"form_order"`
	IsDefault    bool                           `json:"is_default"`
	IsBattleOnly bool                           `json:"is_battle_only"`
	IsMega       bool                           `json:"is_mega"`
	FormName     string                         `json:"form_name"`
	Pokemon      NamedAPIResource[Pokemon]      `json:"pokemon"`
	Types        []PokemonFormType              `json:"types"`
	Sprites      PokemonFormSprites             `json:"pokemon_form_sprites"`
	VersionGroup NamedAPIResource[VersionGroup] `json:"version_group"`
	Names        []Name                         `json:"names"`
	FormNames    []Name                         `json:"form_names"`
}

type PokemonFormSprites struct {
	FrontDefault string `json:"front_default"`
	FrontShiny   string `json:"front_shiny"`
	BackDefault  string `json:"back_default"`
	BackShiny    string `json:"back_shiny"`
}

type PokemonHabitat struct {
	ID             int                              `json:"id"`
	Name           string                           `json:"name"`
	Names          []Name                           `json:"names"`
	PokemonSpecies NamedAPIResource[PokemonSpecies] `json:"pokemon_species"`
}

type PokemonShape struct {
	ID             int                                `json:"id"`
	Name           string                             `json:"name"`
	AwesomeNames   []AwesomeName                      `json:"awesome_names"`
	Names          []Name                             `json:"names"`
	PokemonSpecies []NamedAPIResource[PokemonSpecies] `json:"pokemon_species"`
}

type AwesomeName struct {
	AwesomeName string                     `json:"awesome_name"`
	Language    NamedAPIResource[Language] `json:"language"`
}

type PokemonSpecies struct {
	ID                   int                              `json:"id"`
	Name                 string                           `json:"name"`
	Order                int                              `json:"order"`
	GenderRate           int                              `json:"gender_rate"`
	CaptureRate          int                              `json:"capture_rate"`
	BaseHappiness        int                              `json:"base_happiness"`
	IsBaby               bool                             `json:"is_baby"`
	IsLegendary          bool                             `json:"is_legendary"`
	IsMythical           bool                             `json:"is_mythical"`
	HatchCounter         int                              `json:"hatch_counter"`
	HasGenderDifferences bool                             `json:"has_gender_differences"`
	FormsSwitchable      bool                             `json:"forms_switchable"`
	GrowthRate           NamedAPIResource[GrowthRate]     `json:"growth_rate"`
	PokedexNumders       []PokemonSpeciesDexEntry         `json:"pokedex_numbers"`
	EggGroups            []NamedAPIResource[EggGroup]     `json:"egg_groups"`
	Color                NamedAPIResource[PokemonColor]   `json:"color"`
	Shape                NamedAPIResource[PokemonShape]   `json:"shape"`
	EvolvesFromSpecies   NamedAPIResource[PokemonSpecies] `json:"evolves_from_species"`
	EvolutionChain       APIResource[EvolutionChain]      `json:"evolution_chain"`
	Habitat              NamedAPIResource[PokemonHabitat] `json:"habitat"`
	Generation           NamedAPIResource[Generation]     `json:"generation"`
	Names                []Name                           `json:"names"`
	PalParkEncounters    []PalParkEncounterArea           `json:"pal_park_encounters"`
	FlavorTextEntries    []FlavorText                     `json:"flavor_text_entries"`
	FormDescriptions     []Description                    `json:"form_descriptions"`
	Genera               []Genus                          `json:"genera"`
	Varieties            []PokemonSpeciesVariety          `json:"varieties"`
}

type Genus struct {
	Genus    string                     `json:"genus"`
	Language NamedAPIResource[Language] `json:"language"`
}

type PokemonSpeciesDexEntry struct {
	EntryNumber int                       `json:"entry_number"`
	Pokedex     NamedAPIResource[Pokedex] `json:"pokedex"`
}

type PalParkEncounterArea struct {
	BaseScore int                           `json:"base_score"`
	Rate      int                           `json:"rate"`
	Area      NamedAPIResource[PalParkArea] `json:"area"`
}

type PokemonSpeciesVariety struct {
	IsDefault bool                      `json:"is_default"`
	Pokemon   NamedAPIResource[Pokemon] `json:"pokemon"`
}

type Stat struct {
	ID               int                               `json:"id"`
	Name             string                            `json:"name"`
	GameIndex        int                               `json:"game_index"`
	IsBattleOnly     bool                              `json:"is_battle_only"`
	AffectingMoves   MoveStatAffectSets                `json:"affecting_moves"`
	AffectingNatures NatureStatAffectSets              `json:"affecting_natures"`
	Characteristics  []APIResource[Characteristic]     `json:"characteristics"`
	MoveDamageClass  NamedAPIResource[MoveDamageClass] `json:"move_damage_class"`
	Names            []Name                            `json:"names"`
}

type MoveStatAffectSets struct {
	Increase []MoveStatAffect `json:"increase"`
	Decrease []MoveStatAffect `json:"decrease"`
}

type MoveStatAffect struct {
	Change int                    `json:"change"`
	Move   NamedAPIResource[Move] `json:"move"`
}

type NatureStatAffectSets struct {
	Increase []NamedAPIResource[Nature] `json:"increase"`
	Decrease []NamedAPIResource[Nature] `json:"decrease"`
}

type Type struct {
	ID                  int                               `json:"id"`
	Name                string                            `json:"string"`
	DamageRelations     TypeRelations                     `json:"damage_relations"`
	PastDamageRelations []TypeRelationsPast               `json:"past_damage_relations"`
	GameIndices         []GenerationGameIndex             `json:"game_indices"`
	Generation          NamedAPIResource[Generation]      `json:"generation"`
	MoveDamageClass     NamedAPIResource[MoveDamageClass] `json:"move_damage_class"`
	Names               []Name                            `json:"names"`
	Pokemon             []TypePokemon                     `json:"pokemon"`
	Moves               []NamedAPIResource[Move]          `json:"moves"`
}

type TypePokemon struct {
	Slot    int                       `json:"slot"`
	Pokemon NamedAPIResource[Pokemon] `json:"pokemon"`
}

type TypeRelations struct {
	NoDamageTo       []NamedAPIResource[Type] `json:"no_damage_to"`
	HalfDamageTo     []NamedAPIResource[Type] `json:"half_damage_to"`
	DoubleDamageTo   []NamedAPIResource[Type] `json:"double_damage_to"`
	NoDamageFrom     []NamedAPIResource[Type] `json:"no_damage_from"`
	HalfDamageFrom   []NamedAPIResource[Type] `json:"half_damage_from"`
	DoubleDamageFrom []NamedAPIResource[Type] `json:"double_damage_from"`
}

type TypeRelationsPast struct {
	Generation      NamedAPIResource[Generation] `json:"generation"`
	DamageRelations TypeRelations                `json:"damage_relations"`
}
