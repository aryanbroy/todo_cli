package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

type TodoContent struct {
	Name string `json:"name"`
}

func main() {
	task := flag.String("add", "[Empty]", "Add a task")

	flag.Parse()

	nameJson, err := json.MarshalIndent(TodoContent{Name: *task}, "", " ")

	if err != nil {
		panic(err)
	}

	file, err := os.Create("todos.json")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	_, err = file.Write(nameJson)

	if err != nil {
		panic(err)
	}

	fmt.Println("Task added: ", *task)
	fmt.Println(string(nameJson))
}
