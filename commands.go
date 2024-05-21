package main

import (
	"errors"
	"strconv"
)

type Command interface {
	Execute(args []string, jsonPath string) bool
}

var task Task

type Add struct{}

type Delete struct{}

func (command Add) Execute(args []string, jsonPath string) bool {
	return false
	// err, task := checkTaskAdd()

	// tasks := readTasks(jsonPath)

	// tasks = append(tasks, task)
	// json, err := json.MarshalIndent(tasks, "", "	")

	// if checkErr(err) {
	// 	return false
	// }

	// err = os.WriteFile(jsonPath, json, 0644)

	// return !checkErr(err)
}

// operation / task / priority
func checkTaskAdd(args []string) (bool, error) {
	task.Task = args[1]

	if len(args) == 3 {
		priority, err := strconv.Atoi(args[2])
		if checkErr(err) {
			return false, errors.New("Error while adding priority")
		}
		task.Priority = priority
	}

	return true, nil
}

// func (command Delete) Execute(args []string, jsonPath string) bool {
// 	return false // Pending
// }
