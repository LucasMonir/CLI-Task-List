package main

type Command interface {
	Execute() bool
}

type Add struct {
	Name string
}

type Delete struct {
	Name string
}

func (command Add) Execute() bool {
	return false // Pending
}

func (command Delete) Execute() bool {
	return false // Pending
}
