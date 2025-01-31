package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Todo struct {
	Todos []TodoContent
}

type TodoContent struct {
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

func main() {
	// task := flag.String("add", "[Empty]", "Add a task")
	// done := flag.Bool("done", false, "Mark a task")

	// flag.Parse()

	// nameJson, err := json.MarshalIndent(TodoContent{Name: *task, Done: *done}, "", " ")

	// if err != nil {
	// 	panic(err)
	// }

	// file, err := os.Create("todos.json")

	// defer file.Close()

	var todo []TodoContent 

	todo, err := readExistingData("todos.json")

	if err != nil {
		panic(err)
	}

	fmt.Println(todo[0])

	// _, err = file.Write(nameJson)

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("Task added: ", *task)
	// fmt.Println(string(nameJson))
}
