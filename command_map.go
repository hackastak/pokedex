package main

import (
	"errors"
	"fmt"
)

func commandMapF(cfg *config) error {
	locationsResponse, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationsResponse.Next
	cfg.prevLocationsURL = locationsResponse.Previous

	for _, loc := range locationsResponse.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func commandMapB(cfg *config) error {
	if cfg.prevLocationsURL == nil {
		return errors.New("You're on the first page already")
	}

	locationsResponse, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationsResponse.Next
	cfg.prevLocationsURL = locationsResponse.Previous

	for _, loc := range locationsResponse.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
