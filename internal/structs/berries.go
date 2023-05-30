package structs

type Berry struct {
	ID               int                             `json:"id"`
	Name             string                          `json:"name"`
	GrowthTime       int                             `json:"growth_time"`
	MaxHarvest       int                             `json:"max_harvest"`
	NaturalGiftPower int                             `json:"natural_gift_power"`
	Size             int                             `json:"size"`
	Smoothness       int                             `json:"smoothness"`
	SoilDryness      int                             `json:"soil_dryness"`
	Firmness         NamedAPIResource[BerryFirmness] `json:"firmness"`
	Flavors          []BerryFlavorMap                `json:"flavors"`
	Item             NamedAPIResource[Item]          `json:"item"`
	NaturalGiftType  NamedAPIResource[Type]          `json:"natural_gift_type"`
}

type BerryFlavorMap struct {
	Potency int                           `json:"potency"`
	Flavor  NamedAPIResource[BerryFlavor] `json:"flavor"`
}

type BerryFirmness struct {
	ID      int                       `json:"id"`
	Name    string                    `json:"name"`
	Berries []NamedAPIResource[Berry] `json:"berries"`
	Names   []Name                    `json:"names"`
}

type BerryFlavor struct {
	ID          int                           `json:"id"`
	Name        string                        `json:"name"`
	Berries     []FlavorBerryMap              `json:"berries"`
	ContestType NamedAPIResource[ContestType] `json:"contest_type"`
	Names       []Name                        `json:"names"`
}

type FlavorBerryMap struct {
	Potency int                     `json:"potency"`
	Berry   NamedAPIResource[Berry] `json:"berry"`
}
