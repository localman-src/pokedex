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

type endPoints struct {
	LocationArea string
}

var EndPoints = endPoints{
	LocationArea: "/location-area",
}

func GetResourceList(endpoint string, offset, limit int) (areaList *structs.NamedAPIResourceList[any], err error) {
	url := fmt.Sprintf("%s%s?limit=%d&offset=%d", baseURL, endpoint, limit, offset)

	body, exists := cache.Get(url)
	if exists {
		fmt.Println("Taking Cached Value", url)
		err = json.Unmarshal(body, &areaList)
		if err != nil {
			fmt.Printf("error unmarshalling cached JSML: %v\n", err)
			fmt.Println("attempting to request from pokeapi.co")
		} else {
			return areaList, err
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

	err = json.Unmarshal(body, &areaList)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling JSML: %v", err)
	}

	return areaList, err
}
