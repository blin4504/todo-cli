package main

import (
	"fmt"
	"os"
)

type Todo struct {
	id       int
	task     string
	complete bool
}

var todoList []Todo

func parseCommand(args []string) {
	command := args[0]
	switch command {
	case "add":
		fmt.Println("add commmnd")
	case "list":
		fmt.Println("list command")
	case "remove":
		fmt.Println("remove command")
	case "complete":
		fmt.Println("complete command")
	}
}

func main() {
	args := os.Args

	if len(args) < 3 {
		fmt.Println("Usage: ./todo-cli <command> <args>")
		return
	}

	parseCommand(args[1:])
}
