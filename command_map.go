package main

import (
	"errors"
	"fmt"
)

func commandMapNext(cfg *config, args ...string) error {
	locationsList, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}
	cfg.nextLocationsURL = locationsList.Next
	cfg.prevLocationsURL = locationsList.Previous

	for _, loc := range locationsList.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapPrev(cfg *config, args ...string) error {
	if cfg.prevLocationsURL == nil {
		return errors.New("you are on the first page")
	}

	locationsList, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationsList.Next
	cfg.prevLocationsURL = locationsList.Previous

	for _, loc := range locationsList.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
