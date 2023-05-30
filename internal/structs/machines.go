package structs

type Machine struct {
	ID           int                            `json:"id"`
	Item         NamedAPIResource[Item]         `json:"item"`
	Move         NamedAPIResource[Move]         `json:"move"`
	VersionGroup NamedAPIResource[VersionGroup] `json:"version_group"`
}
