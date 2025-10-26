package main

import (
	"bufio"
	"os"
	"fmt"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

func main() {
	commands := getCommands()
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Pokedex > ")
	for scanner.Scan() {
		input := scanner.Text()
		words := cleanInput(input)
		
		if len(words) == 0 {
			fmt.Print("Pokedex > ")
			continue
		}

		commandName := words[0]
		command, exists := commands[commandName]
		
		if !exists {
			fmt.Printf("Unknown command: %s\n", commandName)
			fmt.Print("Pokedex > ")
			continue
		}

		err := command.callback()
		if err != nil {
			fmt.Printf("Error executing command: %s\n", err)
		}

		fmt.Print("Pokedex > ")
	}
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	
	commands := getCommands()
	for _, cmd := range commands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func cleanInput(text string) []string {
	//The purpose of this function will be to split the user's input into "words" based on whitespace. 
	// It should also lowercase the input and trim any leading or trailing whitespace.
	text = strings.TrimSpace(text)
	text = strings.ToLower(text)
	
	words := strings.Fields(text)
	return words
}