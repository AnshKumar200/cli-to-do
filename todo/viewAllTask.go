package todo

import (
	"fmt"

	"github.com/AnshKumar200/cli-to-do/models"
)

func ViewAllTask(TodoData []models.TodoItem) {
	fmt.Println("-PENDING-")
	for i, val := range TodoData {
		if !val.Completed {
			fmt.Println("Index:", i, " Task:", val.Task)
		}
	}
	fmt.Println("-COMPLETED-")
	for i, val := range TodoData {
		if val.Completed {
			fmt.Println("Index:", i, " Task:", val.Task)
		}
	}
}
