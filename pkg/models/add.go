package models

import (
	util "clitest/pkg/utilitary"
	"encoding/json"
	"fmt"
)

type Add struct{}

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
	var task = Task{}
	_, err := checkTaskAdd(args)

	if util.CheckErr(err) {
		return false
	}

	tasks := ReadTasks()

	if len(tasks) != 0 {
		task.Id = tasks[len(tasks)-1].Id + 1
	} else {
		task.Id = 1
	}

	task.Task = args[1]

	tasks = append(tasks, task)
	json, err := json.MarshalIndent(tasks, "", "	")

	if util.CheckErr(err) {
		return false
	}

	err = util.WriteJson(json)

	if util.CheckErr(err) {
		fmt.Println("Error while creating task, aborting...")
		return false
	}

	fmt.Printf("Task '%d - %s' created sucessfully!", task.Id, task.Task)
	return true
}
