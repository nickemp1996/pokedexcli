package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
	"pokedexcli/internal/config"
)

func cleanInput(text string) []string {
	strings := strings.Fields(strings.ToLower(text))

	return strings
}

func run(reg map[string]cliCommand) {
	cfg := config.Config{}

	scanner := bufio.NewScanner(os.Stdin)

	for ;;{
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			input := cleanInput(scanner.Text())
			command, ok := reg[input[0]]
			if ok {
				err := command.callback(&cfg, input[1:])
				if err != nil {
					fmt.Println("Error running command '", command.name, "':", err)
				}
			} else {
				fmt.Println("Unknown command")
			}

		}

		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading input:", err)
			break;
		}	
	}
}