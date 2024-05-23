package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Path struct {
	TasksPath string `json:"tasksPath"`
}

var invalidTaskSymbols = "!\"#$%&'()*+;/=?@[\\]^_{|}~-"

func GetTaskFilePath() string {
	configsJson, err := os.ReadFile("configs.json")

	if CheckErr(err) {
		return ""
	}

	var path Path

	err = json.Unmarshal(configsJson, &path)

	if CheckErr(err) {
		return ""
	}

	return path.TasksPath
}

func CheckErr(err error) bool {
	if err != nil {
		fmt.Println(err)
		return true
	}
	return false
}

func CheckArgs(args []string) bool {
	return len(args) != 0
}

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
