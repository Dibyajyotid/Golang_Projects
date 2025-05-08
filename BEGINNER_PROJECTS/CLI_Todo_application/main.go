package main

import "fmt"

func main() {
	fmt.Println("CLI Todo Application")

	todos := Todos{}

	storage := NewStorage[Todos]("todos.json")
	storage.Load(&todos)

	cmdFlags := NewCmd()
	cmdFlags.Execute(&todos)
	
	storage.Save(todos)
}