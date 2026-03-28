package main

import (
	"fmt"
	"time"
)

func menu() {
	fmt.Println("Welcome to the Task Manager!")
	fmt.Println("Please select an option:")
	fmt.Println("1. Add Task")
	fmt.Println("2. View Tasks")
	fmt.Println("3. Complete Task")
	fmt.Println("4. Exit")
	var input string
	fmt.Scanln(&input)
	// while
	for input != "4" {
		switch input {
		case "1":
			fmt.Println("Adding a new task...")
		case "2":
			fmt.Println("Viewing tasks...")
		case "3":
			fmt.Println("Completing a task...")
		default:
			fmt.Println("Invalid option, please try again.")
		}
		fmt.Println("Please select an option:")
		fmt.Scanln(&input)
	}

}

var tasks []Task

// Task represents a to-do item with relevant details such as name, completion status, due date, and creation time.
type Task struct {
	ID        int
	Name      string
	Completed bool
	DueDate   time.Time
	CreatedAt time.Time
}

func createTask(name string, dueDate time.Time) Task {
	task := Task{
		ID:        generateID(),
		Name:      name,
		Completed: false,
		DueDate:   dueDate,
		CreatedAt: time.Now(),
	}
	// todo: add task to the list of tasks
	tasks = append(tasks, task)

}

func generateID() int {
	return int(time.Now().UnixNano())
}
