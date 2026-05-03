package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"task-cli/storage"
)

func main() {
	if len(os.Args) < 2 {
		printHelp()
		return
	}
	command := os.Args[1]
	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Usage: task-cli and <task name>")
			return
		}
		taskName := strings.Join(os.Args[2:], " ")
		addTask(taskName)
	case "list":
		listTask()
	case "done":
		if len(os.Args) < 3 {
			fmt.Println("Usage: task-cli done <task name>")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid ID - must be a number")
			return
		}
		markDone(id)
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Usage: task-cli delete <task id>")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid ID - must be a number ")
			return
		}
		deleteTask(id)
	default:
		printHelp()
	}
}

func addTask(name string) {
	tasks, err := storage.LoadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}
	newTask := storage.Task{
		ID:   len(tasks) + 1,
		Name: name,
		Done: false,
	}

	tasks = append(tasks, newTask)

	err = storage.SaveTasks(tasks)
	if err != nil {
		fmt.Println("Error saving task:", err)
		return
	}
	fmt.Println("Task added: [%d] %s\n", newTask.ID, newTask.Name)
}
func listTask() {
	tasks, err := storage.LoadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}
	if len(tasks) == 0 {
		fmt.Println("No tasks yet. Add one with: task-cli add <task name>")
		return
	}

	fmt.Println("\n Your Tasks:")
	fmt.Println("----------------------")

	for _, task := range tasks {
		status := "[ ]"
		if task.Done {
			status = "[✓]"
		}
		fmt.Printf("%s %d. %s\n", status, task.ID, task.Name)
	}
	fmt.Println("----------------------")
}
func markDone(id int) {
	tasks, err := storage.LoadTasks()
	if err != nil {
		fmt.Println("Error loaing tasks:", err)
		return
	}
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Done = true
			storage.SaveTasks(tasks)
			fmt.Println("Marked done: %s\n", task.Name)
			return
		}
	}
	fmt.Printf("No task found with ID %d\n", id)
}
func deleteTask(id int) {
	tasks, err := storage.LoadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			storage.SaveTasks(tasks)
			fmt.Printf("Deleted: %s\n", task.Name)
			return
		}
	}
	fmt.Printf("No task found with ID %d\n,id")
}

func printHelp() {
	fmt.Println(`
Task cli - Manage your tasks from the terminal

commands:
add <name>				Add a new task
list					List all tasks
done <id>				Mark a task as done
delete <id>				Delete a task

`)
}

var _ = bufio.NewReader
