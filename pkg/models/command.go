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

func (command Add) Name() string {
	return "add"
}

func (command Add) ArgCount() int {
	return 2
}

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

func (command List) Name() string {
	return "ls"
}

func (command List) Execute(_ []string) bool {
	tasks := ReadTasks()

	PrintTasks(tasks)

	return true
}

func (command List) ArgCount() int {
	return 1
}

func (Command Delete) Name() string {
	return "del"
}

func (command Delete) ArgCount() int {
	return 2
}

func CheckCommandParams(command Command, args int) bool {
	return command.ArgCount() == args
}

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

func findItemIndex(tasks []Task, id int) (int, error) {
	for index, task := range tasks {
		if task.Id == id {
			return index, nil
		}
	}

	return 0, fmt.Errorf("No item found matching %d", id)
}

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
