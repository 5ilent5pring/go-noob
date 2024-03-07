package main

import (
	"fmt"
	"os"
)

// Define a struct to represent a task
type Task struct {
	Description string
	Completed   bool
}

func main() {
	todoList := []Task{} // Create an empty todo list

	for {
		fmt.Println("\nTodo List Manager")
		fmt.Println("1. Add Task")
		fmt.Println("2. View Tasks")
		fmt.Println("3. Mark Task as Completed")
		fmt.Println("4. Exit")
		fmt.Print("Enter your choice: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			addTask(&todoList)
		case 2:
			viewTasks(todoList)
		case 3:
			markTaskAsCompleted(&todoList)
		case 4:
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

// Function to add a task to the todo list
func addTask(todoList *[]Task) {
	var description string
	fmt.Print("Enter task description: ")
	fmt.Scanln(&description)

	task := Task{Description: description, Completed: false}
	*todoList = append(*todoList, task)

	fmt.Println("Task added successfully.")
}

// Function to view all tasks in the todo list
func viewTasks(todoList []Task) {
	if len(todoList) == 0 {
		fmt.Println("Todo list is empty.")
		return
	}

	fmt.Println("\nTodo List:")
	for i, task := range todoList {
		fmt.Printf("%d. %s - %s\n", i+1, task.Description, getStatus(task.Completed))
	}
}

// Function to mark a task as completed
func markTaskAsCompleted(todoList *[]Task) {
	viewTasks(*todoList)

	var index int
	fmt.Print("Enter the index of the task to mark as completed: ")
	fmt.Scanln(&index)

	if index < 1 || index > len(*todoList) {
		fmt.Println("Invalid index.")
		return
	}

	(*todoList)[index-1].Completed = true
	fmt.Println("Task marked as completed.")
}

// Function to get the status of a task
func getStatus(completed bool) string {
	if completed {
		return "Completed"
	} else {
		return "Pending"
	}
}
