package main

import "fmt"

type Task struct {
	Task     string `json:"task"`
	Id       int    `json:"id"`
	Priority int    `json:"priority"`
}

func (task *Task) ToString() string {
	return fmt.Sprintf("ID: %d | Task: %v | Priority: %d", task.Id, task.Task, task.Priority)
}
