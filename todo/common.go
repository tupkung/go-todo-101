package todo

func NewList() *todoList {
	return &todoList{}
}

func NewTodo(title string) *todoModel {
	return &todoModel{
		title: title,
	}
}
