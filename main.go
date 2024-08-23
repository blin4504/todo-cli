package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Todo struct {
	id       int
	task     string
	complete bool
}

func (todo Todo) String() string {
	return fmt.Sprintf("%d %s %t", todo.id, todo.task, todo.complete)
}

var todoList map[int]Todo = make(map[int]Todo)

func add(task []string) {
	id := len(todoList) + 1
	newTodo := Todo{id, strings.Join(task, " "), false}
	todoList[id] = newTodo
}

func list() {
	for _, todo := range todoList {
		fmt.Println(todo)
	}
}

func remove(id int) {
	delete(todoList, id)
}

func complete(id int) {
	old := todoList[id]
	todoList[id] = Todo{id, old.task, true}
}

func parseCommand(args []string) {
	command := args[0]
	switch command {
	case "add": // add <task>
		add(args[1:])
	case "list": // list
		list()
	case "remove": // remove <id>
		id, err := strconv.Atoi(args[1])
		if err != nil {
			panic(err)
		}
		remove(id)
	case "complete": // complete <id>
		id, err := strconv.Atoi(args[1])
		if err != nil {
			panic(err)
		}
		complete(id)
	case "quit":
		os.Exit(1)
	default:
		fmt.Println("Usage: todo-cli > <command> <args>")
	}
}

func main() {
	for {
		fmt.Print("todo-cli > ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		inp := scanner.Text()
		tokens := strings.Split(inp, " ")
		parseCommand(tokens)
	}
}
