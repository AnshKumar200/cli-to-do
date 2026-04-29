package models

type TodoItem struct {
	Task      string `json:"task"`
	Completed bool   `json:"completed"`
}
