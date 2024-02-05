package main

import (
	"bufio"
	"fmt"
	"os"
)

type CliCommand struct {
	name        string
	description string
	callback    func()
}

var dexCommands = map[string]CliCommand{
	"help": {
		name:        "help",
		description: "Display a help message",
		callback:    commandHelp,
	},
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	},
}

func commandHelp() {
	fmt.Println("Help function was called")
}

func commandExit() {
	fmt.Println("Exit function was called")
}

// func checkCommandExists(line string) CliCommand {
// 	if _, ok
// }

func main() {
	fmt.Println("Welcome to the Hackastak Pokedex! Use 'help' for a list of available commands.")
	buf := bufio.NewReader(os.Stdin)
	fmt.Println("Pokedex> ")

	cmd, err := buf.ReadBytes('\n')
	if err != nil {
		fmt.Println(err)
	} else {
		cmdString := string(cmd)
	}

}
