package main

import (
	"clitest/pkg/models"
	utils "clitest/pkg/utilitary"
	"fmt"
	"os"
	"slices"
	"strings"
)

var availableCommands []string

var commands = map[string]models.Command{
	models.Command.Name(models.Add{}):  models.Add{},
	models.Command.Name(models.List{}): models.List{},
}

func main() {
	buildCommands()

	args := os.Args[1:]

	if !utils.CheckArgs(args) || !checkCommand(args[0]) {
		fmt.Println("Command not found, available commands: ", strings.Join(availableCommands, ", "))
		return
	}

	command := commands[args[0]]

	if !models.CheckCommandParams(command, len(args)) {
		fmt.Printf("Incompatible number of arguments, expected: %d", command.ArgCount())
		return
	}

	command.Execute(args)
}

// buildCommands builds the available commands list
func buildCommands() {
	for command := range commands {
		availableCommands = append(availableCommands, command)
	}
}

// checkCommand checks if command inserted is contained in the available commands list
func checkCommand(command string) bool {
	return slices.Contains(availableCommands[:], command)
}
