package main

import (
	"fmt"
)

func addTask(args []string) {
	if len(args) == 0 {
		fmt.Println("Please provide a task title.")
		return
	}

	title := ""
	for _, a := range args {
		title += a + " "
	}

	//Prevent duplicate tasks
	for _, t := range tasks {
		if t.Title == title {
			fmt.Println("Task already exists:", title)
			return
		}
	}

	tasks = append(tasks, Task{Title: title, Done: false})
	fmt.Println("Added task:", title)
}

func editTask(args []string) {
	if len(args) < 2 {
		fmt.Println("Provide the task number and a new title.")
		return
	}

	i := parseIndex(args[0])
	if i == -1 || args[1] == "" {
		return
	}

	tasks[i].Title = args[1]
	fmt.Println("Edited task:", tasks[i].Title)
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

func markUndone(args []string) {
	if len(args) == 0 {
		fmt.Println("Provide the task number.")
		return
	}

	i := parseIndex(args[0])
	if i == -1 {
		return
	}

	tasks[i].Done = false
	fmt.Println("Marked as undone:", tasks[i].Title)
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
