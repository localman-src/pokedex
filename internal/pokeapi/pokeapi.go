package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/localman-src/pokedex/internal/structs"
)

func PrintResourceList(resourceList structs.NamedAPIResourceList) {
	for _, location := range resourceList.Results {
		println(location.Name)
	}
}
func GetAreas(url string) (*structs.NamedAPIResourceList, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error fetching areas: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}
	locationResponse := structs.NamedAPIResourceList{}

	err = json.Unmarshal(body, &locationResponse)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling JSML: %v", err)
	}

	return &locationResponse, nil
}
