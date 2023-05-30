package structs

type Machine struct {
	ID           int              `json:"id"`
	Item         NamedAPIResource `json:"item"`
	Move         NamedAPIResource `json:"move"`
	VersionGroup NamedAPIResource `json:"version_group"`
}
