package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

var availableCommands = [3]string{"add", "rm", "ls"}

func main() {
	args := os.Args[1:]
	if checkArgs(args) {
		fmt.Println("Task created!")
	}
}

// 1:command 2: task 3: modifier
func checkArgs(args []string) bool {
	if len(args) < 2 {
		fmt.Println("Usage: <Command> <taskName>")
		return false
	}

	if len(args) > 3 {
		fmt.Println("No command uses 3 arguments...")
		return false
	}

	if !checkCommand(args[0]) {
		fmt.Printf("Invalid Command, available commands: %v", availableCommands)
		return false
	}

	if !checkTask(args[1]) {
		fmt.Println("Invalid task name...")
	}

	return true
}

func checkCommand(command string) bool {
	return slices.Contains(availableCommands[:], command)
}

func checkTask(task string) bool {
	trimmedTask := strings.TrimSpace(task)

	if len(trimmedTask) == 0 {
		return false
	}

	if strings.ContainsAny(trimmedTask, "-,_+=)(.;[]}{\\|/?!@#$%¨&*_^^~~'\"") {
		return false
	}

	return true
}
