package main

import (
	"fmt"
	"go-todo-101/todo"
)

func main() {
	list := todo.NewList()
	// add 3 todo
	list.AddTodo(todo.NewTodo("task 1"))
	list.AddTodo(todo.NewTodo("task 2"))
	list.AddTodo(todo.NewTodo("task 3"))

	fmt.Println(list)

	// complete 1 todo
	completeTodo := list.GetTodo(1)
	completeTodo.Complete()

	fmt.Println(list)

	// remove complete todos
	list.RemoveComplete()

	fmt.Println(list)
}
