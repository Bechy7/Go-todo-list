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
	case "list":
		listTasks()
	case "done":
		markDone(os.Args[2:])
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

func addTask(args []string) {
	if len(args) == 0 {
		fmt.Println("Please provide a task title.")
		return
	}

	title := ""
	for _, a := range args {
		title += a + " "
	}

	tasks = append(tasks, Task{Title: title, Done: false})
	fmt.Println("Added task:", title)
}

func listTasks() {
	if len(tasks) == 0 {
		fmt.Println("No tasks yet.")
		return
	}

	for i, t := range tasks {
		status := " "
		if t.Done {
			status = "✓"
		}
		fmt.Printf("%d. [%s] %s\n", i, status, t.Title)
	}
}

func markDone(args []string) {
	if len(args) == 0 {
		fmt.Println("Provide the task number.")
		return
	}

	i := parseIndex(args[0])
	if i == -1 {
		return
	}

	tasks[i].Done = true
	fmt.Println("Marked as done:", tasks[i].Title)
}

func deleteTask(args []string) {
	if len(args) == 0 {
		fmt.Println("Provide the task number.")
		return
	}

	i := parseIndex(args[0])
	if i == -1 {
		return
	}

	fmt.Println("Deleted:", tasks[i].Title)
	tasks = append(tasks[:i], tasks[i+1:]...)
}

func parseIndex(s string) int {
	var idx int
	_, err := fmt.Sscanf(s, "%d", &idx)
	if err != nil || idx < 0 || idx >= len(tasks) {
		fmt.Println("Invalid task number.")
		return -1
	}
	return idx
}
