package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/macedonsky777/pokedexcli/internal/pokeapi"
	"github.com/macedonsky777/pokedexcli/internal/pokecache"
)

var appConfig config

type config struct {
	NextURL       *string
	PreviousURL   *string
	Cache         *pokecache.Cache
	CaughtPokemon map[string]pokeapi.PokemonStruct
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func commandExit(config *config, args []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(config *config, commands map[string]cliCommand, args []string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for _, command := range commands {
		fmt.Printf("%s %s\n", command.name, command.description)
	}
	return nil
}

func getCommands() map[string]cliCommand {
	commands := make(map[string]cliCommand)

	commands["exit"] = cliCommand{
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	}

	commands["help"] = cliCommand{
		name:        "help",
		description: "description of Pokedex commands",
		callback: func(config *config, args []string) error {
			return commandHelp(config, commands, args)
		},
	}
	commands["map"] = cliCommand{
		name:        "map",
		description: "Show next 20 locations",
		callback:    commandMap,
	}
	commands["mapb"] = cliCommand{
		name:        "mapb",
		description: "Show previous 20 locations",
		callback:    commandMapb,
	}
	commands["explore"] = cliCommand{
		name:        "explore",
		description: "show all pokemons in chosen location",
		callback:    commandExplore,
	}
	commands["catch"] = cliCommand{
		name:        "catch",
		description: "Try to catch them all!",
		callback:    commandCatch,
	}
	commands["inspect"] = cliCommand{
		name:        "inspect",
		description: "Show Pokemon stats",
		callback:    commandInspect,
	}
	commands["pokedex"] = cliCommand{
		name:        "pokedex",
		description: "Show all caught pokemon",
		callback:    commandPokedex,
	}
	return commands
}

func commandMap(config *config, args []string) error {
	var url string
	if config.NextURL == nil {
		url = pokeapi.BaseURL
	} else {
		url = *config.NextURL
	}
	if cachedData, exists := config.Cache.Get(url); exists {
		fmt.Println("Кеш отримано!")
		var result pokeapi.MainDataStruct
		if err := json.Unmarshal(cachedData, &result); err == nil {
			for _, location := range result.Results {
				fmt.Println(location.Name)
			}
			config.NextURL = result.Next
			config.PreviousURL = result.Previous
			return nil
		}
	}
	data, err := pokeapi.GetLocationAreas(url)
	if err != nil {
		fmt.Println("Щось пішло не так, локацій немає", err)
		return err
	}
	if jsonData, err := json.Marshal(data); err == nil {
		config.Cache.Add(url, jsonData)
	}
	for _, location := range data.Results {
		fmt.Println(location.Name)
	}
	config.NextURL = data.Next
	config.PreviousURL = data.Previous

	return nil
}

func commandMapb(config *config, args []string) error {
	var url string
	if config.PreviousURL != nil {
		url = *config.PreviousURL
	} else {
		fmt.Println("you're on the first page")
		return nil
	}
	data, err := pokeapi.GetLocationAreas(url)
	if err != nil {
		fmt.Println("Щось пішло не так...", err)
	} else {
		for _, location := range data.Results {
			fmt.Println(location.Name)
		}
		config.NextURL = data.Next
		config.PreviousURL = data.Previous
	}
	return nil
}

func commandExplore(config *config, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("explore потребує назву локації")
	}
	locationName := args[0]
	url := "https://pokeapi.co/api/v2/location-area/" + locationName
	if cachedData, exists := config.Cache.Get(url); exists {
		fmt.Println("Кеш отримано!")
		var result pokeapi.LocationDetailStruct
		if err := json.Unmarshal(cachedData, &result); err == nil {
			fmt.Println("Exploring..." + locationName)
			fmt.Println("Found Pokemon:")
			for _, location := range result.PokemonEncounters {
				fmt.Printf("- %s\n", location.Pokemon.Name)
			}
			return nil
		}
	}
	data, err := pokeapi.GetLocationArea(locationName)
	if err != nil {
		fmt.Println("Чи локації такої нема, чи ти дурачок...", err)
		return err
	}
	if jsonData, err := json.Marshal(data); err == nil {
		config.Cache.Add(url, jsonData)
	}
	fmt.Println("Exploring..." + locationName)
	fmt.Println("Found Pokemon:")
	for _, encounter := range data.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}
	return nil
}

func commandInspect(config *config, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("inspect потребує назви покемона")
	}
	pokemonName := args[0]

	pokemon, exists := config.CaughtPokemon[pokemonName]
	if !exists {
		fmt.Println("you have not caught that pokemon")
		return nil
	}
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Printf("Stats:\n")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Printf("Types:\n")
	for _, PokemonType := range pokemon.Types {
		fmt.Printf("  - %s\n", PokemonType.Type.Name)
	}
	return nil
}

func commandPokedex(config *config, args []string) error {
	if len(args) > 0 {
		return fmt.Errorf("pokedex не потребує додаткових аргументів")
	}
	fmt.Printf("Your Pokedex:\n")
	if len(config.CaughtPokemon) == 0 {
		fmt.Println("No pokemons yet! Explore areas and try to catch someone!")
	} else {
		for pokemonName := range config.CaughtPokemon {
			fmt.Printf(" - %s\n", pokemonName)
		}
	}
	return nil
}

func commandCatch(config *config, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("catch потребує назви покемона")
	}
	pokemonName := args[0]

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
	url := "https://pokeapi.co/api/v2/pokemon/" + pokemonName
	var pokemon pokeapi.PokemonStruct
	if cachedData, exists := config.Cache.Get(url); exists {
		fmt.Println("Кеш отримано")
		json.Unmarshal(cachedData, &pokemon)
	} else {
		data, err := pokeapi.GetPokemon(pokemonName)
		if err != nil {
			return err
		}
		pokemon = data
		if jsonData, err := json.Marshal(pokemon); err == nil {
			config.Cache.Add(url, jsonData)
		}
	}
	pokemonChance := pokemon.BaseExperience
	chance := 100 - pokemonChance
	if rand.Intn(100) < chance {
		fmt.Printf("%s was caught!\n", pokemonName)
		config.CaughtPokemon[pokemonName] = pokemon
		fmt.Println("You may now inspect it with the inspect command.")
	} else {
		fmt.Printf("%s escaped!\n", pokemonName)
	}
	return nil
}

type cliCommand struct {
	name        string
	description string
	callback    func(config *config, args []string) error
}

func main() {
	rand.Seed(time.Now().UnixNano())
	cache := pokecache.NewCache(5 * time.Minute)
	appConfig.Cache = cache
	appConfig.CaughtPokemon = make(map[string]pokeapi.PokemonStruct)
	commands := getCommands()

	for {
		fmt.Print("Pokedex > ")
		scanner := bufio.NewScanner(os.Stdin)
		if !scanner.Scan() {
			break
		}
		userInput := scanner.Text()
		words := cleanInput(userInput)
		if len(words) > 0 {
			command := words[0]
			args := words[1:]
			if command, exists := commands[command]; exists {
				err := command.callback(&appConfig, args)
				if err != nil {
					fmt.Println("Error:", err)
				}
			} else {
				fmt.Println("Unknown command")
				commandExit(&appConfig, args)
			}
		}

	}
}
