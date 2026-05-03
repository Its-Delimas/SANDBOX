package models

import (
	"encoding/json"
	"os"
	"sync"
)

type Task struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Done bool   `json:"done"`
}

const FileName = "tasks.json"

var mu sync.Mutex

func LoadTasks() ([]Task, error) {
	mu.Lock()
	defer mu.Unlock()

	file, err := os.ReadFile(FileName)
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
func SaveTasks(tasks []Task) error {
	mu.Lock()
	defer mu.Unlock()

	data, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile(FileName, data, 0644)
}
