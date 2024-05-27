package commands

import models "clitest/pkg/models"

type List struct{}

// Name returns the name of the command, which is "ls" (list).
func (command List) Name() string {
	return "ls"
}

// Execute executes the "list" command. It reads tasks from a JSON file and prints them.
func (command List) Execute(_ []string) bool {
	tasks := ReadTasks()

	models.PrintTasks(tasks)

	return true
}

// ArgCount returns the required number of arguments for the "list" command, which is 1.
func (command List) ArgCount() int {
	return 1
}
