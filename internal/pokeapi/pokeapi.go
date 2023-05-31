package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/localman-src/pokedex/internal/pokecache"
	"github.com/localman-src/pokedex/internal/structs"
)

const baseURL = "https://pokeapi.co/api/v2"

var cache = pokecache.NewCache()

func GetResourceList(endpoint string, offset, limit int) (resourceList *structs.NamedAPIResourceList[any], err error) {
	url := fmt.Sprintf("%s%s?limit=%d&offset=%d", baseURL, endpoint, limit, offset)

	body, exists := cache.Get(url)
	if exists {
		err = json.Unmarshal(body, &resourceList)
		if err != nil {
			fmt.Printf("error unmarshalling cached data: %v\n", err)
			fmt.Println("attempting to request from pokeapi.co")
		} else {
			cache.Set(url)
			return resourceList, err
		}
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error fetching areas: %v", err)
	}
	defer resp.Body.Close()

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	fmt.Printf("Adding %s data to cache.\n", url)
	cache.Add(url, body)

	err = json.Unmarshal(body, &resourceList)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling JSON: %v", err)
	}

	return resourceList, err
}

func GetLocationArea(name string) (locationAreaObject *structs.LocationArea, err error) {
	resource := structs.APIResource[structs.LocationArea]{
		URL: fmt.Sprintf("%s/location-area/%s/", baseURL, name),
	}

	locationAreaObject, err = CachedGet(resource)

	return locationAreaObject, err
}

func GetPokemon(name string) (pokemonObject *structs.Pokemon, err error) {
	resource := structs.APIResource[structs.Pokemon]{
		URL: fmt.Sprintf("%s/pokemon/%s/", baseURL, name),
	}

	pokemonObject, err = CachedGet(resource)
	return pokemonObject, err
}

func GetSpecies(name string) (pokemonObject *structs.PokemonSpecies, err error) {
	resource := structs.APIResource[structs.PokemonSpecies]{
		URL: fmt.Sprintf("%s/pokemon/%s/", baseURL, name),
	}

	pokemonObject, err = CachedGet(resource)
	return pokemonObject, err
}

func CachedGet[T structs.PokeAPIObject](resource structs.APIResource[T]) (object *T, err error) {
	body, exists := cache.Get(resource.URL)
	if exists {
		err := json.Unmarshal(body, &object)
		if err != nil {
			fmt.Printf("error unmarshalling cached data: %v\n", err)
			fmt.Println("attempting to request from pokeapi.co")
		} else {
			cache.Set(resource.URL)
			return object, err
		}
	}

	object, err = resource.GetPokeAPI()
	return object, err
}

func NamedAPIResourceToAPIResource[T structs.PokeAPIObject](resource structs.NamedAPIResource[T]) structs.APIResource[T] {
	return structs.APIResource[T]{
		URL: resource.URL,
	}
}
