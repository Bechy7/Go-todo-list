package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Task struct {
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

var tasks []Task

func main() {
	loadTasks()

	if len(os.Args) < 2 {
		fmt.Println("Commands: add <task>, list, done <index>, delete <index>")
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		addTask(os.Args[2:])
	case "edit":
		editTask(os.Args[2:])
	case "list":
		listTasks()
	case "done":
		markDone(os.Args[2:])
	case "undone":
		markUndone(os.Args[2:])
	case "delete":
		deleteTask(os.Args[2:])
	default:
		fmt.Println("Unknown command")
	}

	saveTasks()
}

func loadTasks() {
	file, err := os.ReadFile("tasks.json")
	if err != nil {
		return
	}
	json.Unmarshal(file, &tasks)
}

func saveTasks() {
	data, _ := json.MarshalIndent(tasks, "", "  ")
	os.WriteFile("tasks.json", data, 0644)
}
