package main

import (
	"encoding/json"
	"errors"
	"os"
	"strconv"
)

type Command interface {
	Execute(args []string, jsonPath string) bool
	Name() string
	ArgCount() int
}

var task Task

type Add struct{}
type List struct{}

// Add
func (command Add) Name() string {
	return "add"
}

func (command Add) ArgCount() int {
	return 2
}

func (command Add) Execute(args []string, jsonPath string) bool {
	_, err := checkTaskAdd(args)

	if checkErr(err) {
		return false
	}

	tasks := readTasks(jsonPath)
	task.Id = tasks[len(tasks)-1].Id
	tasks = append(tasks, task)
	json, err := json.MarshalIndent(tasks, "", "	")

	if checkErr(err) {
		return false
	}

	err = os.WriteFile(jsonPath, json, 0644)

	return !checkErr(err)
}

// List
func (command List) Name() string {
	return "ls"
}

func (command List) Execute(_ []string, jsonPath string) bool {
	tasks := readTasks(jsonPath)

	printTasks(tasks)

	return true
}

func (command List) ArgCount() int {
	return 1
}

// Generic
func checkTaskAdd(args []string) (bool, error) {
	checkTask(args[1])
	task.Task = args[1]

	if len(args) == 3 {
		priority, err := strconv.Atoi(args[2])
		if checkErr(err) {
			return false, errors.New("error while adding priority")
		}
		task.Priority = priority
	}

	return true, nil
}
