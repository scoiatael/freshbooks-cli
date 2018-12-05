package main

import (
	"log"
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
	projectList, err := freshbooks.ListProjects()
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}

	taskList, err := freshbooks.ListTasks()
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}

	Present(projectList, taskList)
}
