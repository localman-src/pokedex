package structs

type EncounterMethod struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Order string `json:"order"`
	Names []Name `json:"names"`
}

type EncounterCondition struct {
	ID     int                                         `json:"id"`
	Name   string                                      `json:"name"`
	Names  []Name                                      `json:"names"`
	Values []NamedAPIResource[EncounterConditionValue] `json:"values"`
}

type EncounterConditionValue struct {
	ID        int                                  `json:"id"`
	Name      string                               `json:"name"`
	Condition NamedAPIResource[EncounterCondition] `json:"condition"`
	Names     []Name                               `json:"names"`
}
