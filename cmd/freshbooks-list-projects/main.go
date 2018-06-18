package main

import (
	"encoding/xml"
	"fmt"

	"github.com/scoiatael/gofreshbooks"
)

func Present(projectList freshbooks.ProjectList, taskList freshbooks.TaskList) {
	fmt.Printf("Projects:\n")
	for _, project := range projectList.Projects {
		fmt.Printf("- %s (%s)\n", project.Name, project.ID)
	}
	fmt.Printf("Tasks:\n")
	for _, task := range taskList.Tasks {
		fmt.Printf("- %s (%s)\n", task.Name, task.ID)
	}
}

func main() {
	projectListBytes, err := freshbooks.Do(freshbooks.Request{
		Method: "project.list",
	})
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	projectList := freshbooks.ProjectList{}
	if err := xml.Unmarshal(projectListBytes, &projectList); err != nil {
		fmt.Printf("error: %v\n", err)
	}

	taskListBytes, err := freshbooks.Do(freshbooks.Request{
		Method: "task.list",
	})
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	taskList := freshbooks.TaskList{}
	if err := xml.Unmarshal(taskListBytes, &taskList); err != nil {
		fmt.Printf("error: %v\n", err)
	}

	Present(projectList, taskList)
}
