package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Cmd struct {
	Add    string
	Del    int
	Edit   string
	Toggle int
	List   bool
}

func NewCmd() *Cmd {
	cf := Cmd{}

	flag.StringVar(&cf.Add, "add", "", "Add a new todo by specify title")
	flag.StringVar(&cf.Edit, "Edit", "", "Edit a todo by index & specify a new tite.")
	flag.IntVar(&cf.Del, "Del", -1, "Delete a todo by index.")
	flag.IntVar(&cf.Toggle, "Toggle", -1, "Toggle a todo by index.")
	flag.BoolVar(&cf.List, "List", false, "List all todos")

	flag.Parse()

	return &cf
}

func (cf *Cmd) Execute(todos *Todos) {
	switch {
	case cf.List:
		todos.print()

	case cf.Add != "":
		todos.add(cf.Add)

	case cf.Edit != "":
		parts := strings.SplitN(cf.Edit, ":", 2)
		if len(parts) != 2 {
			fmt.Println("Error, Invalid format for edit. Please use id:new_title")
			os.Exit(1)
		}

		index, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("Error: Invalid index for edit")
			os.Exit(1)
		}

		todos.edit(index, parts[1])

	case cf.Toggle != -1:
		todos.toggle(cf.Toggle)

	case cf.Del != -1:
		todos.delete(cf.Del)

	default:
		fmt.Println("Invalid command")
	}
}
