package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/0xYotta/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
}

func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)

	for {
		//
		fmt.Print("Pokedex > ")
		reader.Scan()
		//
		userInputFields := cleanInput(reader.Text())
		if reader.Text() == "" {
			continue
		}
		commandName := userInputFields[0]
		args := []string{}
		if len(userInputFields) > 1 {
			args = userInputFields[1:]
		}
		fmt.Printf("Your command was: %s\n", commandName)

		command, exists := getCommands()[commandName]
		//
		if exists {
			err := command.callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}

	}
}

func cleanInput(text string) []string {
	if len(text) == 0 {
		return nil
	}

	return strings.Fields(strings.ToLower(text))
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},

		"explore": {
			name:        "explore <location_name>",
			description: "Explore a location",
			callback:    commandExplore,
		},

		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			callback:    commandMapNext,
		},

		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:    commandMapPrev,
		},
	}
}
