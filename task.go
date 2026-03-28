package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
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
			fmt.Println("Add a new task")
			fmt.Println("Enter name")
			var name string
			fmt.Scanln(&name)
			fmt.Println("Enter due date (YYYY-MM-DD)")
			var dueDateStr string
			fmt.Scanln(&dueDateStr)
			dueDate, err := time.Parse("2006-01-02", dueDateStr)
			if err != nil {
				fmt.Println("Invalid date format. Please use YYYY-MM-DD.")
				break
			}
			createTask(name, dueDate)
			// how to show success message after creating a task?
			fmt.Println("Task created successfully!")

		case "2":
			fmt.Println("Tasks:")
			if len(tasks) == 0 {
				fmt.Println("  (no tasks yet)")
			} else {
				for _, t := range tasks {
					status := "open"
					if t.Completed {
						status = "completed"
					}
					fmt.Printf("- ID: %d; %s; due %s; status: %s\n", t.ID, t.Name, t.DueDate.Format("2006-01-02"), status)
				}
			}
		case "3":
			fmt.Println("Enter ID of task to complete:")
			var idStr string
			fmt.Scanln(&idStr)
			id, err := strconv.ParseInt(idStr, 10, 64)
			if err != nil {
				fmt.Println("Invalid task ID.")
				break
			}
			completed := false
			for i := range tasks {
				if int64(tasks[i].ID) == id {
					tasks[i].Completed = true
					completed = true
					break
				}
			}
			if completed {
				saveTasksToFile()
				fmt.Println("Task marked as completed and saved.")
			} else {
				fmt.Println("Task ID not found.")
			}
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

func createTask(name string, dueDate time.Time) {
	task := Task{
		ID:        generateID(),
		Name:      name,
		Completed: false,
		DueDate:   dueDate,
		CreatedAt: time.Now(),
	}
	// todo: add task to the list of tasks
	tasks = append(tasks, task)
	// add list to tasks.json
	saveTasksToFile()

}

func generateID() int {
	return int(time.Now().UnixNano())
}

func saveTasksToFile() {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		fmt.Println("Error serializing tasks:", err)
		return
	}
	if err := os.WriteFile("tasks.json", data, 0o644); err != nil {
		fmt.Println("Error writing tasks.json:", err)
	}
}

func loadTasksFromFile() {
	if _, err := os.Stat("tasks.json"); err != nil {
		if os.IsNotExist(err) {
			return
		}
		fmt.Println("Error checking tasks.json:", err)
		return
	}
	data, err := os.ReadFile("tasks.json")
	if err != nil {
		fmt.Println("Error reading tasks.json:", err)
		return
	}
	if len(data) == 0 {
		return
	}
	if err := json.Unmarshal(data, &tasks); err != nil {
		fmt.Println("Error parsing tasks.json:", err)
	}
}

func init() {
	loadTasksFromFile()
}
