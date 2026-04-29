package todo

import (
	"encoding/json"
	"fmt"
	"os"
	"slices"

	"github.com/AnshKumar200/cli-to-do/models"
)

func DeleteTask(index int, TodoData *[]models.TodoItem) {
	*TodoData = slices.Delete((*TodoData), index, index+1)
	updaedTask, _ := json.MarshalIndent(*TodoData, "", "  ")
	os.WriteFile("data.json", updaedTask, 0644)
	fmt.Println(index, " is removed!")
}
