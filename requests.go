package freshbooks

import (
	"encoding/xml"
	"log"
)

func ListProjects() (projectList ProjectList, err error) {
	projectListBytes, err := Do(Request{
		Method: "project.list",
	})
	err = xml.Unmarshal(projectListBytes, &projectList)
	if err != nil {
		log.Println(string(projectListBytes))
	}
	return
}

func ListTasks() (taskList TaskList, err error) {
	taskListBytes, err := Do(Request{
		Method: "task.list",
	})
	err = xml.Unmarshal(taskListBytes, &taskList)
	if err != nil {
		log.Println(string(taskListBytes))
	}
	return
}
