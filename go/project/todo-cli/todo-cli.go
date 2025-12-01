package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

// Terminal color codes
const (
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Reset  = "\033[0m"
)

// Todo struct defines a task
type Todo struct {
	ID        int       `json:"id"`
	Task      string    `json:"task"`
	Category  string    `json:"category"`
	DueDate   string    `json:"due_date"`
	Done      bool      `json:"done"`
	CreatedAt time.Time `json:"created_at"`
}

var todos []Todo
var dataFile = "todo_advanced.json"

// ------------------ FUNCTIONS ------------------

// Load todos from JSON file
func loadTodos() {
	file, err := os.Open(dataFile)
	if err != nil {
		if os.IsNotExist(err) {
			todos = []Todo{}
			return
		}
		panic(err)
	}
	defer file.Close()
	json.NewDecoder(file).Decode(&todos)
}

// Save todos to JSON file
func saveTodos() {
	file, err := os.Create(dataFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	json.NewEncoder(file).Encode(todos)
}

// Add a new task
func addTask(task, category, due string) {
	id := len(todos) + 1
	todo := Todo{
		ID:        id,
		Task:      task,
		Category:  category,
		DueDate:   due,
		Done:      false,
		CreatedAt: time.Now(),
	}
	todos = append(todos, todo)
	fmt.Println(Green+"Added task:"+Reset, task)
}

// List all tasks with colors
func listTasks() {
	if len(todos) == 0 {
		fmt.Println(Yellow + "No tasks found." + Reset)
		return
	}
	fmt.Println(Blue + "Your Tasks:" + Reset)
	for _, t := range todos {
		status := Red + "❌" + Reset
		if t.Done {
			status = Green + "✔️" + Reset
		}
		fmt.Printf("%d. [%s] %s (Category: %s, Due: %s)\n", t.ID, status, t.Task, t.Category, t.DueDate)
	}
}

// Mark a task as done
func markDone(id int) {
	for i, t := range todos {
		if t.ID == id {
			todos[i].Done = true
			fmt.Println(Green+"Marked as done:"+Reset, t.Task)
			return
		}
	}
	fmt.Println(Red + "Task not found!" + Reset)
}

// Delete a task
func deleteTask(id int) {
	for i, t := range todos {
		if t.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			fmt.Println(Red+"Deleted:"+Reset, t.Task)
			return
		}
	}
	fmt.Println(Red + "Task not found!" + Reset)
}

// ------------------ MAIN ------------------

func main() {
	loadTodos()

	for {
		fmt.Println("\n" + Blue + "Todo CLI Advanced Menu:" + Reset)
		fmt.Println("1. Add Task")
		fmt.Println("2. List Tasks")
		fmt.Println("3. Mark Done")
		fmt.Println("4. Delete Task")
		fmt.Println("5. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var task, category, due string
			fmt.Print("Enter task: ")
			fmt.Scanln()
			task, _ = readLine()
			fmt.Print("Enter category (Work/Personal/Study): ")
			category, _ = readLine()
			fmt.Print("Enter due date (YYYY-MM-DD): ")
			due, _ = readLine()
			addTask(task, category, due)
		case 2:
			listTasks()
		case 3:
			var id int
			fmt.Print("Enter task ID to mark done: ")
			fmt.Scan(&id)
			markDone(id)
		case 4:
			var id int
			fmt.Print("Enter task ID to delete: ")
			fmt.Scan(&id)
			deleteTask(id)
		case 5:
			saveTodos()
			fmt.Println(Green + "Goodbye!" + Reset)
			return
		default:
			fmt.Println(Red + "Invalid choice, try again!" + Reset)
		}
	}
}

// readLine(): Helper to read multi-word input
func readLine() (string, error) {
	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		buf := make([]byte, 1024)
		n, _ := os.Stdin.Read(buf)
		input = strings.TrimSpace(string(buf[:n]))
	}
	return input, err
}

/*
------------------- EXPLANATION -------------------

1. Terminal Colors
   - Red, Green, Blue, Yellow, Reset for colored output

2. Todo struct
   - Added Category and DueDate fields for advanced functionality

3. addTask()
   - Accepts task name, category, and due date
   - Appends to todos slice

4. listTasks()
   - Shows task status with colored checkmarks
   - Prints category and due date

5. readLine()
   - Reads multi-word input from user
   - Ensures tasks like "Buy milk and bread" work properly

6. Persistent Storage
   - loadTodos() and saveTodos() handle JSON file
   - Tasks are saved between program runs

7. Menu
   - Infinite loop
   - User chooses option
   - Handles Add, List, Mark Done, Delete, Exit

---------------------------------------------------
*/
