package main

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
)

func readTasks() []Task {
	var tasks []Task
	jsonPath := getTaskFilePath()

	if !checkTaskFileExists() {
		initTaskFile()
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

func checkTaskFileExists() bool {
	jsonPath := getTaskFilePath()

	info, err := os.Stat(jsonPath)

	if os.IsNotExist(err) || checkErr(err) || info.IsDir() {
		fmt.Println("Creating task file")
		initTaskFile()
	}

	return true
}

func initTaskFile() {
	jsonPath := getTaskFilePath()

	err := os.WriteFile(jsonPath, []byte(""), fs.ModePerm)
	if checkErr(err) {
		fmt.Println(err.Error())
	}
}
