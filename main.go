package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"
	"time"
)

type Todo struct {
	Name        string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

func createTodo(todoName string, todos *[]Todo) {
	todo := Todo{
		Name:        todoName,
		Completed:   false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}

	*todos = append(*todos, todo)
}

func deleteTodo(todoIndex int, todos *[]Todo) {
	*todos = append((*todos)[:todoIndex], (*todos)[todoIndex+1:]...)
}

func completeTodo(todoIndex int, todos *[]Todo) {
	(*todos)[todoIndex].Completed = true
	(*todos)[todoIndex].CompletedAt = time.Now()
}

func listTodos(todos *[]Todo) {
	for index, todo := range *todos {
		fmt.Printf("Task id: %d - Name: %s - Completed: %t - Created at: %d/%d/%d - Completed at: %d/%d/%d \n", index, todo.Name, todo.Completed, todo.CreatedAt.Year(), todo.CreatedAt.Month(), todo.CreatedAt.Day(), todo.CreatedAt.Year(), todo.CreatedAt.Month(), todo.CreatedAt.Day())
	}
}

func main() {
	add := flag.String("a", "", "Add a todo")
	delete := flag.Int("d", -1, "Delete a todo")
	complete := flag.Int("c", -1, "Complete a todo")
	list := flag.Bool("l", false, "List all todos")

	flag.Parse()

	var todos []Todo

	data, err := os.ReadFile("todos.json")
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			os.Create("todos.json")
		}
	}

	json.Unmarshal(data, &todos)

	switch {
	case len(*add) > 0:
		createTodo(*add, &todos)
		fmt.Println("Task", *add, "created!")
	case *delete > -1:
		deleteTodo(*delete, &todos)
		fmt.Println("Task", *delete, "deleted!")
	case *complete > -1:
		completeTodo(*complete, &todos)
		fmt.Println("Task", *complete, "completed!")
	case *list:
		listTodos(&todos)
	}

	data, _ = json.Marshal(todos)
	os.WriteFile("todos.json", data, 0755)
}
