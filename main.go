package main

import (
	"os"
)

var jsonPath = "task.json"
var availableCommands = [3]string{"add", "rm", "ls"}
var invalidTaskSymbols = "!\"#$%&'()*+;/=?@[\\]^_{|}~-"

func main() {
	args := os.Args[1:]
	if !checkArgs(args) {
		return
	}

	tasks := readTasks(jsonPath)
	printTasks(tasks)
}
