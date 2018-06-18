package freshbooks

import (
	"encoding/xml"
	"time"
)

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

type TimeEntry struct {
	ProjectId string  `xml:"project_id"`
	TaskId    string  `xml:"task_id"`
	Hours     float64 `xml:"hours"`
	Notes     string  `xml:"notes"`
	Date      Date    `xml:"date"`
}

type TimeEntryId struct {
	XMLName xml.Name `xml:"response"`
	Id      string   `xml:"time_entry_id"`
}

type Date struct {
	time.Time
}

func (date Date) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeElement(date.Time.Format("2006-01-02"), start)
	return nil
}
