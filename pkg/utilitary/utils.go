// Package Utils is dedicated to general utilities used by other files
package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// Path is used to unmarshal the path of the task.json file
type Path struct {
	TasksPath string `json:"tasksPath"`
}

var invalidTaskSymbols = "!\"#$%&'()*+;/=?@[\\]^_{|}~-"

// GetTaskFilePath recovers the path of the task.json from the configuration file
func GetTaskFilePath() string {
	configsJson, err := os.ReadFile("../../configs.json")

	if CheckErr(err) {
		return ""
	}

	var path Path

	err = json.Unmarshal(configsJson, &path)

	if CheckErr(err) {
		return ""
	}

	return "../../" + path.TasksPath
}

// CheckErr is utilized to check if an error happened and to log it to console.
func CheckErr(err error) bool {
	if err != nil {
		fmt.Println(err)
		return true
	}
	return false
}

// CheckArgs checks length of args param
func CheckArgs(args []string) bool {
	return len(args) != 0
}

// CheckTask validates if task already exists in the listing
func CheckTask(task string) bool {
	trimmedTask := strings.TrimSpace(task)

	if len(trimmedTask) == 0 {
		return false
	}

	if strings.ContainsAny(trimmedTask, invalidTaskSymbols) {
		return false
	}

	return true
}
