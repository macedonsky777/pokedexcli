package pokeapi

type MainDataStruct struct {
	Count    int
	Next     *string
	Previous *string
	Results  []LocationDataStruct
}

type LocationDataStruct struct {
	Name string
	URL  string
}

type LocationDetailStruct struct {
	PokemonEncounters []PokemonEncounter `json:"pokemon_encounters"`
}

type PokemonEncounter struct {
	Pokemon PokemonInfo `json:"pokemon"`
}

type PokemonInfo struct {
	Name string `json:"name"`
}

type PokemonStruct struct {
	Name           string        `json:"name"`
	BaseExperience int           `json:"base_experience"`
	Height         int           `json:"height"`
	Weight         int           `json:"weight"`
	Stats          []PokemonStat `json:"stats"`
	Types          []PokemonType `json:"types"`
}

type PokemonStat struct {
	BaseStat int      `json:"base_stat"`
	Stat     StatInfo `json:"stat"`
}

type StatInfo struct {
	Name string `json:"name"`
}

type PokemonType struct {
	Type TypeInfo `json:"type"`
}

type TypeInfo struct {
	Name string `json:"name"`
}
