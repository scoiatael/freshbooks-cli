package freshbooks

import (
	"encoding/xml"
)

func ListProjects() (projectList ProjectList, err error) {
	projectListBytes, err := Do(Request{
		Method: "project.list",
	})
	err = xml.Unmarshal(projectListBytes, &projectList)
	return
}

func ListTasks() (taskList TaskList, err error) {
	taskListBytes, err := Do(Request{
		Method: "task.list",
	})
	err = xml.Unmarshal(taskListBytes, &taskList)
	return
}
