package todo

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/AnshKumar200/cli-to-do/models"
)

func AddTask(task string, TodoData *[]models.TodoItem) {
	newTask := models.TodoItem{Task: task, Completed: false}
	*TodoData = append(*TodoData, newTask)

	updatedTask, _ := json.MarshalIndent(*TodoData, "", "  ")
	os.WriteFile("data.json", updatedTask, 0644)
	fmt.Println("Added!", task)
}
