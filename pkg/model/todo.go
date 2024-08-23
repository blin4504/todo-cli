package model

import "fmt"

type Todo struct {
	ID       int
	Task     string
	Complete bool
}

func (todo Todo) String() string {
	return fmt.Sprintf("%d %s %t", todo.ID, todo.Task, todo.Complete)
}
