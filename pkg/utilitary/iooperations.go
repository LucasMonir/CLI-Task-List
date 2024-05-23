package utils

import (
	"fmt"
	"io/fs"
	"os"
)

func CheckTaskFileExists() bool {
	jsonPath := GetTaskFilePath()

	info, err := os.Stat(jsonPath)

	if os.IsNotExist(err) || CheckErr(err) || info.IsDir() {
		fmt.Println("Creating task file")
		InitTaskFile()
	}

	return true
}

func InitTaskFile() {
	jsonPath := GetTaskFilePath()

	err := os.WriteFile(jsonPath, []byte(""), fs.ModePerm)
	if CheckErr(err) {
		fmt.Println(err.Error())
	}
}

func WriteJson(json []byte) error {
	return os.WriteFile(GetTaskFilePath(), json, 0644)
}
