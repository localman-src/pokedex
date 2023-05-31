package structs

type PokeAPIResource interface {
	GetPokeAPI()
}

type PokeAPIObject interface {
	Pokemon | LocationArea | PokemonSpecies
}
