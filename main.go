package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

// create an auto increment id field
type TodoContent struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Done bool `json:"done"`
}

func readExistingData(filename string) ([]TodoContent, error) {
	data, err := os.ReadFile(filename)

	if err != nil {
		fmt.Println("Error while reading file")
		fmt.Println(err)
		return nil, err
	}

	var todos []TodoContent

	err = json.Unmarshal(data, &todos)

	return todos,err 
}

func saveTodo(filename string, todoData []byte) error  {
	return os.WriteFile(filename, todoData, 0644)
}

func main() {

	task := flag.String("add", "[Empty]", "Add a task")
	done := flag.Bool("done", false, "Mark a task")

	filename := "todos.json"
	flag.Parse()

	taskByte, err := json.MarshalIndent(TodoContent{Name: *task, Done: *done}, "", " ")

	if err != nil {
		panic(err)
	}

	todos, err := readExistingData(filename)

	if err != nil {
		taskByte, error := json.MarshalIndent([]TodoContent{{Id: 1, Name: *task, Done: *done}}, "", " ")
		if error != nil {
			fmt.Println("Error while creating first task")
			return
		}
		os.WriteFile(filename, taskByte, 0644)
		return
	}

	var inputTodo TodoContent
	err = json.Unmarshal(taskByte, &inputTodo)

	if err != nil {
		panic(err)
	}

	// fmt.Println(newTodo)
	// fmt.Println(todos)

	inputTodo.Id = len(todos) + 1

	newTodoSlice := append(todos, inputTodo)
	
	newTodoByte, err := json.MarshalIndent(newTodoSlice, "", " ")

	if err != nil {
		panic(err)
	}

	err = saveTodo(filename, newTodoByte)

	if err != nil {
		panic(err)
	}

}
