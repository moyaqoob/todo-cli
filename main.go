package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Todo struct {
	ID       int
	Title    string
	Complete bool
}

var todos []Todo

func addTodo() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your todo\n")
	todo, err := reader.ReadString('\n')

	newTodo := Todo{
		ID:       len(todos) + 1,
		Title:    strings.TrimSpace(todo),
		Complete: false,
	}

	if err != nil {
		fmt.Println("cannot read the todos")
	}

	todos = append(todos, newTodo)
	fmt.Println("Todo added successfully")
}

func deleteTodo() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter the id of todo you wanna delete")
	var updatedTodos []Todo

	input,_ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	id,err:= strconv.Atoi(input)

	if err!=nil{
		fmt.Println("Invalid ID")
	}

	for _,todo :=range todos{
		if todo.ID != id{
			updatedTodos = append(updatedTodos,todo)
		}
	}
	todos = updatedTodos
	
}

func showTodos() {
	if len(todos) == 0 {
		fmt.Println("No todos yet.")
		return
	}

	fmt.Println("\nAll Todos:")
	for _, t := range todos {
		status := "❌"
		if t.Complete {
			status = "✅"
		}
		fmt.Printf("ID: %d | %s | Done: %s\n", t.ID, t.Title, status)
	}
}

func markTodo() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter the id of todo you wanna mark completed")

	input,_ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	id,err:= strconv.Atoi(input)

	if err!=nil{
		fmt.Println("Invalid id \n")
	}

	for i,todo :=range todos{
		if todo.ID == id{
			todos[i].Complete = true
			fmt.Println("Todo mark as completed")
			return
		}
	}

	fmt.Println("Todo not found")

}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("\nHere are the list of operations \n 1.Add \n 2.Delete \n 3.showAllTodos\n 4.Complete\n 5.quit")
	
	for{
		fmt.Print("\n Enter your choice?\n")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)
		switch choice {
		case "add":
			addTodo()
		case "delete":
			deleteTodo()
		case "ShowAll":
			showTodos()
		case "complete":
			markTodo()
		case "quit":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid command")
		}
	}


}
