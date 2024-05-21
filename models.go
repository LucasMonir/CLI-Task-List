package main

type Task struct {
	Task     string `json:"task"`
	Id       int    `json:"id"`
	Priority int    `json:"priority"`
}

type Tasks struct {
	Tasks []Task `json:"tasks"`
}
