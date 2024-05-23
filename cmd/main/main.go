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

	if !utils.CheckArgs(args) {
		return
	}

	if !CheckCommand(args[0]) {
		fmt.Println("Command not found, available commands: ", strings.Join(availableCommands, ", "))
		return
	}

	command := commands[args[0]]

	if !models.CheckCommandParams(command, len(args)) {
		fmt.Printf("Incompatible number of arguments, expected: %d", command.ArgCount())
	}

	command.Execute(args)
}

func buildCommands() {
	for command := range commands {
		availableCommands = append(availableCommands, command)
	}
}

func CheckCommand(command string) bool {
	return slices.Contains(availableCommands[:], command)
}
