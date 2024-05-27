package models

import (
	util "clitest/pkg/utilitary"
	"encoding/json"
	"fmt"
	"strconv"
)

type Delete struct{}

// Name returns the name of the command, which is "del" (delete).
func (Command Delete) Name() string {
	return "del"
}

// ArgCount returns the required number of arguments for the "delete" command, which is 2.
func (command Delete) ArgCount() int {
	return 2
}

// Execute executes the "delete" command. finds the specified id and remove the task if exists.
func (command Delete) Execute(args []string) bool {
	id, err := strconv.Atoi(args[1])

	if util.CheckErr(err) {
		return false
	}

	tasks := ReadTasks()
	index, err := findItemIndex(tasks, id)
	task := tasks[index]

	if util.CheckErr(err) {
		return false
	}

	tasks = append(tasks[:index], tasks[index+1:]...)
	json, err := json.MarshalIndent(tasks, "", "	")

	if util.CheckErr(err) {
		return false
	}

	err = util.WriteJson(json)

	if util.CheckErr(err) {
		return false
	}

	fmt.Println("Task removed: ", task.ToString())

	return true
}
