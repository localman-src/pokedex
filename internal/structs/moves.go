package structs

type Move struct {
	ID                 int                               `json:"id"`
	Name               string                            `json:"name"`
	Accuracy           int                               `json:"accuracy"`
	EffectChance       int                               `json:"effect_chance"`
	PP                 int                               `json:"pp"`
	Priority           int                               `json:"priority"`
	Power              int                               `json:"power"`
	ContestCombos      ContestComboSets                  `json:"contest_combos"`
	ContestType        NamedAPIResource[ContestType]     `json:"contest_type"`
	ContestEffect      APIResource[ContestEffect]        `json:"contest_effect"`
	DamageClass        NamedAPIResource[MoveDamageClass] `json:"damage_class"`
	EffectEntries      []VerboseEffect                   `json:"effect_entries"`
	EfectChanges       []AbilityEffectChange             `json:"effect_changes"`
	LearnedByPokemon   []NamedAPIResource[Pokemon]       `json:"learned_by_pokemon"`
	FlavorTextEntries  []MoveFlavorText                  `json:"flavor_text_entries"`
	Generation         NamedAPIResource[Generation]      `json:"generation"`
	Machines           []MachineVersionDetail            `json:"machines"`
	Meta               MoveMetaData                      `json:"meta"`
	Names              []Name                            `json:"names"`
	PastValues         []PastMoveStatValues              `json:"past_values"`
	StatChanges        []MoveStatChange                  `json:"stat_changes"`
	SuperContestEffect APIResource[SuperContestEffect]   `json:"super_contest_effect"`
	Target             NamedAPIResource[MoveTarget]      `json:"target"`
	Type               NamedAPIResource[Type]            `json:"type"`
}

type ContestComboSets struct {
	Normal ContestComboDetail `json:"normal"`
	Super  ContestComboDetail `json:"super"`
}

type ContestComboDetail struct {
	UseBefore []NamedAPIResource[Move] `json:"use_before"`
	UseAfter  []NamedAPIResource[Move] `json:"id"`
}

type MoveFlavorText struct {
	FlavorText   string                         `json:"flavor_text"`
	Language     NamedAPIResource[Language]     `json:"language"`
	VersionGroup NamedAPIResource[VersionGroup] `json:"version_group"`
}

type MoveMetaData struct {
	Ailment       NamedAPIResource[MoveAilment]  `json:"ailment"`
	Category      NamedAPIResource[MoveCategory] `json:"category"`
	MinHits       int                            `json:"min_hits"`
	MaxHits       int                            `json:"max_hits"`
	MinTurns      int                            `json:"min_turns"`
	MaxTurns      int                            `json:"max_turns"`
	Drain         int                            `json:"drain"`
	Healing       int                            `json:"healing"`
	CritRate      int                            `json:"crit_rate"`
	AilmentChance int                            `json:"ailment_chance"`
	FlinchChance  int                            `json:"flinch_chance"`
	StatChance    int                            `json:"stat_chance"`
}

type MoveStatChange struct {
	Change int                    `json:"change"`
	Stat   NamedAPIResource[Stat] `json:"stat"`
}

type PastMoveStatValues struct {
	Accuracy      int                            `json:"accuracy"`
	EffectChance  int                            `json:"effect_chance"`
	Power         int                            `json:"power"`
	PP            int                            `json:"pp"`
	EffectEntries []VerboseEffect                `json:"effect_entries"`
	Type          NamedAPIResource[Type]         `json:"type"`
	VersionGroup  NamedAPIResource[VersionGroup] `json:"version_group"`
}

type MoveAilment struct {
	ID    int                    `json:"id"`
	Name  string                 `json:"name"`
	Moves NamedAPIResource[Move] `json:"moves"`
	Names []Name                 `json:"names"`
}

type MoveBattleStyle struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Names []Name `json:"names"`
}

type MoveCategory struct {
	ID           int                      `json:"id"`
	Name         string                   `json:"name"`
	Moves        []NamedAPIResource[Move] `json:"moves"`
	Descriptions []Description            `json:"descriptions"`
}

type MoveDamageClass struct {
	ID           int                      `json:"id"`
	Name         string                   `json:"name"`
	Descriptions []Description            `json:"descriptions"`
	Moves        []NamedAPIResource[Move] `json:"moves"`
	Names        []Name                   `json:"names"`
}

type MoveLearnMethod struct {
	ID            int                              `json:"id"`
	Name          string                           `json:"name"`
	Descriptions  []Description                    `json:"descriptions"`
	Names         []Name                           `json:"names"`
	VersionGroups []NamedAPIResource[VersionGroup] `json:"version_groups"`
}

type MoveTarget struct {
	ID           int                      `json:"id"`
	Name         string                   `json:"name"`
	Descriptions []Description            `json:"descriptions"`
	Moves        []NamedAPIResource[Move] `json:"moves"`
	Names        []Name                   `json:"names"`
}
