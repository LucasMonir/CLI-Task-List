// Package iooperations is dedicated to methods related to I/O
package utils

import (
	"fmt"
	"io/fs"
	"os"
)

// CheckTaskFileExists validates the existance of the task.json file
func CheckTaskFileExists() bool {
	jsonPath := GetTaskFilePath()

	info, err := os.Stat(jsonPath)

	if os.IsNotExist(err) || CheckErr(err) || info.IsDir() || IsFileEmpty(jsonPath) {
		fmt.Println("Creating task file")
		InitTaskFile()
	}

	return true
}

func IsFileEmpty(jsonPath string) bool {
	file, err := os.ReadFile(jsonPath)

	if CheckErr(err) {
		return true
	}

	if len(file) == 0 {
		return true
	}

	return false
}

// InitTaskFile creates the task.json file
func InitTaskFile() {
	jsonPath := GetTaskFilePath()

	err := os.WriteFile(jsonPath, []byte("[]"), fs.ModePerm)
	if CheckErr(err) {
		fmt.Println(err.Error())
	}
}

// WriteJson writes data to the task.json file
func WriteJson(json []byte) error {
	return os.WriteFile(GetTaskFilePath(), json, 0644)
}
