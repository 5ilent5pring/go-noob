package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

// Define a struct to represent a task
type Task struct {
	Description string
	Completed   bool
}

func main() {
	todoList := []Task{} // Create an empty todo list

	currentDate := time.Now().Format("2006-01-02") // Get the current date in yyyy-mm-dd format
	fileName := currentDate + ".txt"              // Generate a filename based on the current date

	loadTasks(fileName, &todoList) // Load tasks from the file

	for {
		fmt.Println("\nTodo List Manager")
		fmt.Println("1. Add Task")
		fmt.Println("2. View Tasks")
		fmt.Println("3. Mark Task as Completed")
		fmt.Println("4. View Today's Output")
		fmt.Println("5. Exit")
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
			viewTodayOutput(fileName, todoList)
		case 5:
			saveTasks(fileName, todoList) // Save tasks to the file before exiting
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

// Function to view today's output
func viewTodayOutput(fileName string, todoList []Task) {
	fmt.Println("\nToday's Todo List:")
	viewTasks(todoList)
}

// Function to load tasks from the file
func loadTasks(fileName string, todoList *[]Task) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("No tasks for today.")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		description := scanner.Text()
		task := Task{Description: description, Completed: false}
		*todoList = append(*todoList, task)
	}
}

// Function to save tasks to the file
func saveTasks(fileName string, todoList []Task) {
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, task := range todoList {
		if task.Completed {
			fmt.Fprintln(writer, task.Description, "- Completed") // Save task as Completed if it's completed
		} else {
			fmt.Fprintln(writer, task.Description)
		}
	}
	writer.Flush()
}
