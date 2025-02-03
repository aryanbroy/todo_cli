package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
)

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

func displayTodo(t table.Writer, todos []TodoContent) {
	for _, item := range todos {
		t.AppendRow(table.Row{item.Id, item.Name, func () string {
			if item.Done {
				return "Done"
			} else {
				return "Pending"
			}
		}()})
	}
	t.Render()
}

func deleteTodo(todos []TodoContent, indexToDel int, filename string) {
	todos = append(todos[ :indexToDel], todos[indexToDel + 1 : ]...)
	byteSlice, err := json.MarshalIndent(todos, "", " ")	
	if err != nil {
		fmt.Println("Error marshalling todos...")
		panic(err)
	}
	err  = saveTodo(filename, byteSlice)
	if err != nil {
		panic(err)
	}

}

func main() {

	task := flag.String("add", "[Empty]", "Add a task")
	done := flag.Bool("done", false, "Mark a task")
	display := flag.Bool("display", false, "Display todo list")
	deleteTaskId := flag.Int("del", -1, "Delete a specific task using its id")

	filename := "todos.json"
	flag.Parse()

	var todos []TodoContent
	todos, err := readExistingData(filename)

	if err != nil {
		fmt.Println("error happening here")
		panic(err)
	}

	if *deleteTaskId > 0 {
		indexToDel := -1
		for i, item := range todos {
			if item.Id == *deleteTaskId {
				indexToDel = i
				break
			}
		}

		if indexToDel == -1 {
			fmt.Println("No such task to delete...")
			return
		}
		
		deleteTodo(todos, indexToDel, filename)	
		return
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"ID", "NAME", "STATUS"})

	if *display {	
		displayTodo(t, todos)	
		return
	}

	taskByte, err := json.MarshalIndent(TodoContent{Name: *task, Done: *done}, "", " ")	

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