package storage

import (
	"encoding/json"
	"os"
)

type Task struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Done bool   `json:"done"`
}

const fileName = "tasks.json"

// loadTasks reads tasks from the JSON file
func LoadTasks() ([]Task, error) {
	file, err := os.ReadFile(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			return []Task{}, nil
		}
		return nil, err
	}

	var tasks []Task
	err = json.Unmarshal(file, &tasks)
	return tasks, err
}

// saveTasks Writes to the JSON file
func SaveTasks(tasks []Task) error {
	data, err := json.Marshal(tasks)
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, data, 0644)
}
