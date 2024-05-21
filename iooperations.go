package main

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
)

func readTasks(jsonPath string) []Task {
	var tasks []Task

	if !checkTaskFileExists(jsonPath) {
		initTaskFile(jsonPath)
		return tasks
	}

	jsonFile, err := os.ReadFile(jsonPath)

	if checkErr(err) {
		return tasks
	}

	err = json.Unmarshal(jsonFile, &tasks)

	if checkErr(err) {
		return tasks
	}

	return tasks
}

func checkTaskFileExists(jsonPath string) bool {
	info, err := os.Stat(jsonPath)

	if os.IsNotExist(err) || checkErr(err) || info.IsDir() {
		fmt.Println("Creating task file")
		initTaskFile(jsonPath)
	}

	return true
}

func initTaskFile(jsonPath string) {
	err := os.WriteFile(jsonPath, []byte(""), fs.ModePerm)
	if checkErr(err) {
		fmt.Println(err.Error())
	}
}
