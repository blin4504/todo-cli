package service

import (
	"fmt"
	"strings"

	"github.com/blin4504/todo-cli/pkg/model"
)

var todoList map[int]model.Todo = make(map[int]model.Todo)

func Add(task []string) {
	id := len(todoList) + 1
	newTodo := model.Todo{ID: id, Task: strings.Join(task, " "), Complete: false}
	todoList[id] = newTodo
}

func List() {
	for _, todo := range todoList {
		fmt.Println(todo)
	}
}

func Remove(id int) {
	delete(todoList, id)
}

func Complete(id int) {
	old := todoList[id]
	todoList[id] = model.Todo{ID: id, Task: old.Task, Complete: true}
}
