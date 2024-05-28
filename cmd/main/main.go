package main

import (
	"fmt"
	"os"
	"slices"
	"strings"

	"clitest/pkg/commands"
	utils "clitest/pkg/utilitary"
)

var commandsMap = map[string]commands.Command{
	commands.Command.Name(commands.Add{}):    commands.Add{},
	commands.Command.Name(commands.List{}):   commands.List{},
	commands.Command.Name(commands.Delete{}): commands.Delete{},
}

func main() {
	availableCommands := buildCommands()

	args := os.Args[1:]

	if err := validateArgs(args, availableCommands); err != nil {
		fmt.Println(err)
		return
	}

	command := commandsMap[args[0]]

	if ok := commands.CheckCommandParams(command, len(args)); !ok {
		fmt.Printf("Incompatible number of arguments, expected: %d\n", command.ArgCount())
		return
	}

	command.Execute(args)
}

// buildCommands builds the available commands list
func buildCommands() []string {
	var commandsList []string
	for command := range commandsMap {
		commandsList = append(commandsList, command)
	}
	return commandsList
}

// validateArgs checks if the arguments are valid and if the command exists
func validateArgs(args, availableCommands []string) error {
	if !utils.CheckArgs(args) || !slices.Contains(availableCommands, args[0]) {
		return fmt.Errorf("command not found, available commands: %s", strings.Join(availableCommands, ", "))
	}
	return nil
}
