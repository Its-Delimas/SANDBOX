package routes

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"task-api/models"
)

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/tasks", tasksHandler)
	mux.HandleFunc("/tasks/", taskHandler)
}

// handle /tasks
func tasksHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getTasks(w, r)
	case http.MethodPost:
		createTask(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// handles /tasks/{id}
func taskHandler(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/tasks/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	switch r.Method {
	case http.MethodPut:
		markDone(w, r, id)
	case http.MethodDelete:
		deleteTask(w, r, id)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := models.LoadTasks()
	if err != nil {
		http.Error(w, "Failed to load tasks", http.StatusInternalServerError)
		return
	}
	respondJSON(w, http.StatusOK, tasks)
}

func createTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	tasks, err := models.LoadTasks()
	if err != nil {
		http.Error(w, "Failed to load tasks", http.StatusInternalServerError)
		return
	}
	task.ID = len(tasks) + 1
	task.Done = false
	tasks = append(tasks, task)

	err = models.SaveTasks(tasks)
	if err != nil {
		http.Error(w, "Failed to save task", http.StatusInternalServerError)
		return
	}
	respondJSON(w, http.StatusCreated, task)
}
func markDone(w http.ResponseWriter, r *http.Request, id int) {
	tasks, err := models.LoadTasks()
	if err != nil {
		http.Error(w, "Failed to load tasks", http.StatusInternalServerError)
		return
	}
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Done = true
			models.SaveTasks(tasks)
			respondJSON(w, http.StatusOK, tasks[i])
			return
		}
	}
	http.Error(w, "Task not found", http.StatusNotFound)
}

func deleteTask(w http.ResponseWriter, r *http.Request, id int) {
	tasks, err := models.LoadTasks()
	if err != nil {
		http.Error(w, "Failed to load tasks", http.StatusInternalServerError)
		return
	}
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:1], tasks[i+1:]...)
			models.SaveTasks(tasks)
			respondJSON(w, http.StatusOK, map[string]string{"message": "Task deleted "})
			return
		}
	}
	http.Error(w, "Task not found", http.StatusNotFound)
}

func respondJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
