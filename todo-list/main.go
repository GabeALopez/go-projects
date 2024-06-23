package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Welcome to the To-Do List Application")

	var tasks []string
	reader := bufio.NewReader(os.Stdin)

	for {
		displayMenu()
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			addTask(&tasks)
		case "2":
			viewTasks(tasks)
		case "3":
			deleteTask(&tasks)
		case "4":
			fmt.Println("Exiting the application.")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}

	}
}

func displayMenu() {
	fmt.Println("\nMenu:")
	fmt.Println("1. Add Task")
	fmt.Println("2. View Tasks")
	fmt.Println("3. Delete Task")
	fmt.Println("4. Exit")
	fmt.Print("Enter your choice: ")
}

func addTask(tasks *[]string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter task description: ")
	task, _ := reader.ReadString('\n')
	task = strings.TrimSpace(task)
	*tasks = append(*tasks, task)
	fmt.Println("Task added!")
}

func viewTasks(tasks []string) {
	fmt.Println("\nYour Tasks:")
	for i, task := range tasks {
		fmt.Printf("%d. %s\n", i+1, task)
	}
}

func deleteTask(tasks *[]string) {
	var taskNum int
	fmt.Print("Enter task number to delete: ")
	fmt.Scan(&taskNum)

	if taskNum > 0 && taskNum <= len(*tasks) {
		*tasks = append((*tasks)[:taskNum-1], (*tasks)[taskNum:]...)
		fmt.Println("Task deleted!")
	} else {
		fmt.Println("Invalid task number!")
	}
}
