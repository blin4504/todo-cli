package ui

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/blin4504/todo-cli/pkg/service"
)

func ParseCommand(args []string) {
	command := args[0]
	switch command {
	case "add":
		service.Add(args[1:])
	case "list":
		service.List()
	case "remove":
		id, err := strconv.Atoi(args[1])
		if err != nil {
			panic(err)
		}
		service.Remove(id)
	case "complete":
		id, err := strconv.Atoi(args[1])
		if err != nil {
			panic(err)
		}
		service.Complete(id)
	case "quit":
		os.Exit(1)
	default:
		fmt.Println("Usage: todo-cli > <command> <args>")
	}
}

func StartCLI() {
	for {
		fmt.Print("todo-cli > ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		inp := scanner.Text()
		tokens := strings.Split(inp, " ")
		ParseCommand(tokens)
	}
}
