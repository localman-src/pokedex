package main

import (
	"fmt"
	"os"

	"github.com/localman-src/pokedex/internal/pokeapi"
)

func commandHelp(cfg *config, params ...string) error {
	if params == nil {
		fmt.Println("Welcome to the Pokedex!")
		cfg.Commands.PrintCommands()
	} else {
		command, err := cfg.Commands.getCommand(params[0])
		if err != nil {
			return err
		}
		fmt.Println(command.helptext)
	}

	return nil
}

func commandExit(cfg *config, params ...string) error {
	os.Exit(0)
	return nil
}

func commandMap(cfg *config, params ...string) error {

	resp, err := pokeapi.GetResourceList("/location-area", cfg.MapOffset, cfg.ResourceLimit)
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return err
	} else {
		resp.Print()
		cfg.MapOffset += cfg.ResourceLimit
		fmt.Printf("New Offset: %d\n", cfg.MapOffset)
	}

	return nil
}

func commandMapB(cfg *config, params ...string) error {

	newOffset := cfg.MapOffset - cfg.ResourceLimit
	if newOffset <= 0 {
		newOffset = 0
	}
	resourceList, err := pokeapi.GetResourceList("/location-area", newOffset, cfg.ResourceLimit)
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return err
	} else {
		resourceList.Print()
		cfg.MapOffset = newOffset
		fmt.Printf("New Offset: %d\n", cfg.MapOffset)
	}
	return nil
}

func commandExplore(cfg *config, params ...string) error {
	locationArea, err := pokeapi.GetLocationArea(params[0])
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return err
	} else {
		locationArea.PrintPokemonEncounters()
	}

	return nil
}

func commandCatch(cfg *config, params ...string) error {
	pokemon, err := pokeapi.GetPokemon(params[0])
	if err != nil {
		fmt.Println(err)
		return err
	}

	species, err := pokeapi.CachedGet(pokemon.Species.ToAPIResource())
	if err != nil {
		fmt.Println(err)
		return err
	}

	catchRoll := cfg.prng.Intn(255)

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	if catchRoll < species.CaptureRate {
		fmt.Printf("Caught: %s (Roll: %d vs Catch Rate: %d)\n", pokemon.Name, catchRoll, species.CaptureRate)
		entry, err := NewPokedexEntry(pokemon)
		if err != nil {
			return err
		}
		cfg.Pokedex.Add(entry)
	} else {
		fmt.Printf("%s escaped! (Roll: %d vs Catch Rate: %d)\n", pokemon.Name, catchRoll, species.CaptureRate)
	}

	return nil
}

func commandInspect(cfg *config, params ...string) error {
	entry, err := cfg.Pokedex.Get(params[0])
	if err != nil {
		return err
	}

	fmt.Printf("Name: %s\n# Caught: %d\nHeight: %d\nWeight: %d\n", entry.Name, entry.Count, entry.Height, entry.Weight)
	fmt.Println("Stats:")
	for _, stat := range entry.Stats {
		fmt.Printf("  - %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Println("Types:")
	for _, poketype := range entry.Types {
		fmt.Printf("  - %s\n", poketype.Type.Name)
	}
	englishFactoids := []string{}
	englishVersions := []string{}
	for _, fact := range entry.Description {
		if fact.Language.Name == "en" {
			englishFactoids = append(englishFactoids, fact.FlavorText)
			englishVersions = append(englishVersions, fact.Version.Name)
		}
	}
	factIndex := cfg.prng.Intn(len(englishFactoids))
	fmt.Printf("Factoid: (%s)\n", englishVersions[factIndex])
	fmt.Println(englishFactoids[factIndex])

	return nil
}

func commandPokedex(cfg *config, params ...string) error {
	fmt.Println("Your Pokedex:")
	for _, entry := range cfg.Pokedex {
		fmt.Printf(" - %dx %s\n", entry.Count, entry.Name)
	}

	return nil
}

func commandSave(cfg *config, params ...string) error {
	err := cfg.Pokedex.Save()

	if err != nil {
		return err
	}

	fmt.Println("Pokedex saved to disk successfully.")

	return nil
}

func commandLoad(cfg *config, params ...string) error {
	loadedDex, err := cfg.Pokedex.Load()
	if err != nil {
		return err
	}

	cfg.Pokedex = *loadedDex

	fmt.Println("Pokedex loaded from disk successfully.")

	return nil
}
