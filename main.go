package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/AnshKumar200/cli-to-do/models"
	"github.com/AnshKumar200/cli-to-do/todo"
)

func main() {
	var exit bool = false
	var choice int

	data, err := os.ReadFile("data.json")
	if err != nil {
		fmt.Println("error: ", err)
		return
	}

	var TodoData []models.TodoItem
	err = json.Unmarshal(data, &TodoData)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}

	for !exit {
		fmt.Println("What do you want to do?")
		fmt.Println("1 - View Tasks")
		fmt.Println("2 - Add")
		fmt.Println("3 - Mark As Complete")
		fmt.Println("4 - Delete")

		fmt.Scan(&choice)

		switch choice {
		case 1:
			todo.ViewAllTask(TodoData)
		case 2:
			var TaskName string
			fmt.Print("Task: ")
			fmt.Scan(&TaskName)
			todo.AddTask(TaskName, &TodoData)
		case 3:
			var index int
			fmt.Println("Enter the index of the task: ")
			todo.MarkTodo(index)
		case 4:
			var index int
			fmt.Println("Enter the index of the task: ")
			todo.DeleteTask(index)
		}

		fmt.Println("=====================")
	}

}
