package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	Id     int    `json:"id"`
	Desc   string `json:"description"`
	IsDone bool   `json:"isDone"`
}

func loadJsonTasks(file string, t *[]Task) {
	file_bytes, err := os.ReadFile(file)
	if err == nil {
		json.Unmarshal(file_bytes, t)
	} else {
		fmt.Println("This file does not exist! Creating a new empty JSON file")
		os.Create("tasks.json")
	}
}

func saveJsonTasks(file string, t []Task) {
	_, err := os.OpenFile(file, os.O_WRONLY, 0644)
	// if file does not exist, create it
	if err != nil {
		fmt.Println("File does not exist, creating now!")
		os.Create("tasks.json")
	}
	taskJson, _ := json.MarshalIndent(t, "", "  ")
	os.WriteFile(file, taskJson, os.FileMode(os.O_WRONLY))
}

func main() {
	// Reads tasks from file and stores them in a slice
	tasks := []Task{}
	loadJsonTasks("tasks.json", &tasks)

	switch os.Args[1] {
	// Displaying all of the tasks
	// Example: go run main.go list
	case "list":
		// if more parameters are passed, inform the user
		if len(os.Args) != 2 {
			fmt.Println("List: Invalid arguments provided with list")
			break
		}
		// if there are no tasks, explicitly tell the user
		if len(tasks) == 0 {
			fmt.Println("No Tasks Present")
			break
		}
		for _, v := range tasks {
			fmt.Println("**************")
			fmt.Println("Task ID:", v.Id)
			fmt.Println("Task Description:", v.Desc)
			fmt.Println("Task Done?", v.IsDone)
		}
	case "add":
		// adds a new task to the list and saves it to json file
		// Example command: "go run main.go add wash dishes"
		if len(os.Args) == 2 {
			// if there are no tasks specified, inform the user
			fmt.Println("No Task added, please enter a valid task description")
			break
		}

		// Create task with input parameters
		t_id := len(tasks) + 1
		t_desc := strings.Join(os.Args[2:], " ")
		t := Task{
			Id:     t_id,
			Desc:   t_desc,
			IsDone: false,
		}
		// add new task to tasks slice
		tasks = append(tasks, t)
		// Save new task to json file
		saveJsonTasks("tasks.json", tasks)
	case "done":
		// Marks the IsDone property of a specific task to true
		// Example Command: "go run main.go done 3" ---> Marks the task with ID 3 as done

		// If an incorrect number of parameters is provided, tell the user
		if len(os.Args) != 3 {
			fmt.Println("Incorrect number of parameters provided.")
			break
		}
		// if a non-int id is provided, tell the user
		value, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("ID provided is not an Integer!.")
			break
		}
		// If a task with the specified ID does not exist, tell the user
		if value > len(tasks) || value <= 0 {
			fmt.Println("This ID does not exist!")
			break
		}
		// If the ID does exist, set it's IsDone property to true
		tasks[value-1].IsDone = true
		// Save the newly updated task to the JSON file
		saveJsonTasks("tasks.json", tasks)
	default:
		// If an incorrect command is used
		fmt.Println("Command Line Argument is not valid!")
	}

}
