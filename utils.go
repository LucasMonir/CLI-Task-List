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
	return len(args) != 0
}

func checkCommandParams(command Command, args int) bool {
	return command.ArgCount() == args
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
		fmt.Println(task.ToString())
	}
}
