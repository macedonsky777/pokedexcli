package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

const BaseURL = "https://pokeapi.co/api/v2/location-area"

func GetLocationAreas(url string) (MainDataStruct, error) {
	response, err := http.Get(url)
	if err != nil {
		return MainDataStruct{}, err
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return MainDataStruct{}, err
	}

	var result MainDataStruct
	err = json.Unmarshal(body, &result)

	if err != nil {
		return MainDataStruct{}, err
	}

	return result, nil

}

func GetLocationArea(locationName string) (LocationDetailStruct, error) {
	url := BaseURL + "/" + locationName
	response, err := http.Get(url)
	if err != nil {
		return LocationDetailStruct{}, err
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return LocationDetailStruct{}, err
	}

	var result LocationDetailStruct

	err = json.Unmarshal(body, &result)

	if err != nil {
		return LocationDetailStruct{}, err
	}

	return result, nil

}

func GetPokemon(pokemonName string) (PokemonStruct, error) {
	url := "https://pokeapi.co/api/v2/pokemon/" + pokemonName
	response, err := http.Get(url)
	if err != nil {
		return PokemonStruct{}, err
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)

	if err != nil {
		return PokemonStruct{}, err
	}

	var result PokemonStruct

	err = json.Unmarshal(body, &result)

	if err != nil {
		return PokemonStruct{}, err
	}
	return result, nil
}
