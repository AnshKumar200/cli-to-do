package todo

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/AnshKumar200/cli-to-do/models"
)

func MarkTodo(index int, TodoData *[]models.TodoItem) {
	(*TodoData)[index].Completed = true
	updatedTask, _ := json.MarshalIndent(*TodoData, "", "  ")
	os.WriteFile("data.json", updatedTask, 0644)
	fmt.Println(index, " is marked as complete!")
}
