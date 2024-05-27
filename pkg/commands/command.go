// Package dedicate to the Command struct and it's operations.
package commands

import (
	models "clitest/pkg/models"
	util "clitest/pkg/utilitary"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
)

type Command interface {
	Execute(args []string) bool
	Name() string
	ArgCount() int
}

// CheckCommandParams checks if the number of arguments matches the required number for a given command.
func CheckCommandParams(command Command, args int) bool {
	return command.ArgCount() == args
}

// checkTaskAdd validates and extracts task details from the arguments provided for adding a task.
func checkTaskAdd(args []string) (bool, error) {
	util.CheckTask(args[1])
	var task = models.Task{}
	task.Task = args[1]

	if len(args) == 3 {
		priority, err := strconv.Atoi(args[2])
		if util.CheckErr(err) {
			return false, errors.New("error while adding priority")
		}
		task.Priority = priority
	}

	return true, nil
}

// findItemIndex finds the index of a task with a given ID in the task list.
func findItemIndex(tasks []models.Task, id int) (int, error) {
	for index, task := range tasks {
		if task.Id == id {
			return index, nil
		}
	}

	return 0, fmt.Errorf("no item found matching given id: %d", id)
}

// ReadTasks reads tasks from a JSON file and returns them as a slice of Task structs.
func ReadTasks() []models.Task {
	var tasks []models.Task
	jsonPath := util.GetTaskFilePath()

	if !util.CheckTaskFileExists() {
		util.InitTaskFile()
		return tasks
	}

	jsonFile, err := os.ReadFile(jsonPath)

	if util.CheckErr(err) {
		return tasks
	}

	err = json.Unmarshal(jsonFile, &tasks)

	if util.CheckErr(err) {
		return tasks
	}

	return tasks
}
