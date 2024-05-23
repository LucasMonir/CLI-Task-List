// Package dedicate to the Command struct and it's operations.
package models

import (
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

type Add struct{}

type List struct{}

type Delete struct{}

var task Task

// Name returns the name of the command, which is "add".
func (command Add) Name() string {
	return "add"
}

// ArgCount returns the required number of arguments for the "add" command, which is 2.
func (command Add) ArgCount() int {
	return 2
}

// Execute executes the "add" command. It reads task details from the arguments, adds the task to the list,
// and writes the updated list to a JSON file.
func (command Add) Execute(args []string) bool {
	_, err := checkTaskAdd(args)

	if util.CheckErr(err) {
		return false
	}

	tasks := ReadTasks()
	task.Id = tasks[len(tasks)-1].Id
	tasks = append(tasks, task)
	json, err := json.MarshalIndent(tasks, "", "	")

	if util.CheckErr(err) {
		return false
	}

	err = util.WriteJson(json)

	return !util.CheckErr(err)
}

// Name returns the name of the command, which is "ls" (list).
func (command List) Name() string {
	return "ls"
}

// Execute executes the "list" command. It reads tasks from a JSON file and prints them.
func (command List) Execute(_ []string) bool {
	tasks := ReadTasks()

	PrintTasks(tasks)

	return true
}

// ArgCount returns the required number of arguments for the "list" command, which is 1.
func (command List) ArgCount() int {
	return 1
}

// Name returns the name of the command, which is "del" (delete).
func (Command Delete) Name() string {
	return "del"
}

// ArgCount returns the required number of arguments for the "delete" command, which is 2.
func (command Delete) ArgCount() int {
	return 2
}

// CheckCommandParams checks if the number of arguments matches the required number for a given command.
func CheckCommandParams(command Command, args int) bool {
	return command.ArgCount() == args
}

// checkTaskAdd validates and extracts task details from the arguments provided for adding a task.
func checkTaskAdd(args []string) (bool, error) {
	util.CheckTask(args[1])
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
func findItemIndex(tasks []Task, id int) (int, error) {
	for index, task := range tasks {
		if task.Id == id {
			return index, nil
		}
	}

	return 0, fmt.Errorf("No item found matching %d", id)
}

// ReadTasks reads tasks from a JSON file and returns them as a slice of Task structs.
func ReadTasks() []Task {
	var tasks []Task
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
