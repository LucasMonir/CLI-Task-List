package main

import (
	commands "clitest/pkg/commands"
	utils "clitest/pkg/utilitary"
	"fmt"
	"os"
	"slices"
	"strings"
)

var availableCommands []string

var commandsMap = map[string]commands.Command{
	commands.Command.Name(commands.Add{}):    commands.Add{},
	commands.Command.Name(commands.List{}):   commands.List{},
	commands.Command.Name(commands.Delete{}): commands.Delete{},
}

func main() {
	buildCommands()

	args := os.Args[1:]

	if !utils.CheckArgs(args) || !checkCommand(args[0]) {
		fmt.Println("Command not found, available commands: ", strings.Join(availableCommands, ", "))
		return
	}

	command := commandsMap[args[0]]

	if !commands.CheckCommandParams(command, len(args)) {
		fmt.Printf("Incompatible number of arguments, expected: %d", command.ArgCount())
		return
	}

	command.Execute(args)

}

// buildCommands builds the available commands list
func buildCommands() {
	for command := range commandsMap {
		availableCommands = append(availableCommands, command)
	}
}

// checkCommand checks if command inserted is contained in the available commands list
func checkCommand(command string) bool {
	return slices.Contains(availableCommands[:], command)
}
