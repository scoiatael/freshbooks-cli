package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Request struct {
	XMLName xml.Name `xml:"request"`
	Method  string   `xml:"method,attr"`
}

type Task struct {
	ID   string `xml:"task_id"`
	Name string `xml:"name"`
}

type Project struct {
	ID    string `xml:"project_id"`
	Name  string `xml:"name"`
	Tasks []Task `xml:"tasks>task"`
}

type ProjectList struct {
	XMLName  xml.Name  `xml:"response"`
	Projects []Project `xml:"projects>project"`
}

type TaskList struct {
	XMLName xml.Name `xml:"response"`
	Tasks   []Task   `xml:"tasks>task"`
}

func Do(request Request) ([]byte, error) {
	api := os.Getenv("FRESHBOOKS_API_URL")
	apiKey := os.Getenv("AUTHENTICATION_TOKEN")

	client := &http.Client{}

	output, err := xml.Marshal(request)
	if err != nil {
		return []byte{}, err
	}

	req, err := http.NewRequest("GET", api, bytes.NewReader(output))
	if err != nil {
		return []byte{}, err
	}
	req.SetBasicAuth(apiKey, "X")

	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func Present(projectList ProjectList, taskList TaskList) {
	fmt.Printf("Projects:\n")
	for _, project := range projectList.Projects {
		fmt.Printf("- %s (%s)\n", project.Name, project.ID)
		for _, taskId := range project.Tasks {
			for _, task := range taskList.Tasks {
				if taskId.ID == task.ID {
					fmt.Printf("  - %s (%s)\n", task.Name, task.ID)
				}
			}
		}
	}
}

func main() {
	projectListBytes, err := Do(Request{
		Method: "project.list",
	})
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	projectList := ProjectList{}
	if err := xml.Unmarshal(projectListBytes, &projectList); err != nil {
		fmt.Printf("error: %v\n", err)
	}

	taskListBytes, err := Do(Request{
		Method: "task.list",
	})
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	taskList := TaskList{}
	if err := xml.Unmarshal(taskListBytes, &taskList); err != nil {
		fmt.Printf("error: %v\n", err)
	}

	Present(projectList, taskList)
}
