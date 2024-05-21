package main

import "os"

var jsonPath = "task.json"
var invalidTaskSymbols = "!\"#$%&'()*+;/=?@[\\]^_{|}~-"
var availableCommands []string

var commands = map[string]Command{
	"add": Add{},
	// "delete": Delete{},
}

func main() {
	buildCommands()
	args := os.Args[1:]

	if !checkArgs(args) {
		return
	}

	command := commands[args[0]]
	command.Execute(args, jsonPath)
	tasks := readTasks(jsonPath)
	printTasks(tasks)
}

func buildCommands() {
	for command := range commands {
		availableCommands = append(availableCommands, command)
	}
}
