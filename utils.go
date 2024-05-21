package main

import (
	"fmt"
	"slices"
	"strings"
)

func checkErr(err error) bool {
	if err != nil {
		fmt.Println(err)
		return true
	}
	return false
}

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
		fmt.Println("Invalid task name, it can't contain the following symbols:", invalidTaskSymbols)
		return false
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

	if strings.ContainsAny(trimmedTask, invalidTaskSymbols) {
		return false
	}

	return true
}

func printTasks(tasks []Task) {
	for _, task := range tasks {
		printTask(task)
	}
}

func printTask(task Task) {
	fmt.Printf("ID: %d", task.Id)
	fmt.Print(" | ")
	fmt.Printf("Task: %v", task.Task)
	fmt.Print(" | ")
	fmt.Printf("Priority: %d", task.Priority)
}
