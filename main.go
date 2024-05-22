package main

import (
	"fmt"
	"os"
	"strings"
)

var invalidTaskSymbols = "!\"#$%&'()*+;/=?@[\\]^_{|}~-"
var availableCommands []string

var commands = map[string]Command{
	Command.Name(Add{}):  Add{},
	Command.Name(List{}): List{},
}

func main() {
	buildCommands()
	args := os.Args[1:]

	if !checkArgs(args) {
		return
	}

	if !checkCommand(args[0]) {
		fmt.Println("Command not found, available commands: ", strings.Join(availableCommands, ", "))
		return
	}

	command := commands[args[0]]

	if !checkCommandParams(command, len(args)) {
		fmt.Printf("Incompatible number of arguments, expected: %d", command.ArgCount())
	}

	command.Execute(args)
}

func buildCommands() {
	for command := range commands {
		availableCommands = append(availableCommands, command)
	}
}
