package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/localman-src/pokedex/internal/structs"
)

func GetAreas(url string) (areaList *structs.NamedAPIResourceList[structs.LocationArea], err error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error fetching areas: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	err = json.Unmarshal(body, &areaList)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling JSML: %v", err)
	}

	return areaList, err
}
