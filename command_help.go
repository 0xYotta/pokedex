package main

import "fmt"

func commandHelp(cfg *config, args ...string) error {
	defer fmt.Println()
	fmt.Printf("\nWelcome to the Pokedex!\nUsage:\n\n")
	commands := getCommands()

	for _, command := range commands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}

	return nil
}
