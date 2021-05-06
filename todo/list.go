package todo

import (
	"fmt"
)

type todoList struct {
	todos []*todoModel
}

func (todoList todoList) GetList() []*todoModel {
	return todoList.todos
}

func (todoList *todoList) AddTodo(tm *todoModel) {
	todoList.todos = append(todoList.todos, tm)
}

func (todoList todoList) GetTodo(index int) *todoModel {
	if len(todoList.todos)-1 >= index {
		return todoList.todos[index]
	}
	return nil
}

func (todoList *todoList) RemoveComplete() {
	newTodos := make([]*todoModel, 0)

	for idx, _ := range todoList.todos {
		if !todoList.todos[idx].isComplete {
			newTodos = append(newTodos, todoList.todos[idx])
		}
	}

	todoList.todos = newTodos
}

func (todoList todoList) String() string {
	result := fmt.Sprintf("total task : %d\n", len(todoList.todos))
	for _, v := range todoList.todos {
		result = fmt.Sprintf("%stitle: %s, isComplete: %v\n", result, v.title, v.isComplete)
	}

	return result
}
