package main

import (
	"encoding/xml"
	"fmt"
	"os"
	"time"

	"github.com/gocarina/gocsv"
	"github.com/scoiatael/gofreshbooks"
)

type Date struct {
	time.Time
}

func (date *Date) MarshalCSV() (string, error) {
	return date.Time.Format("2006-01-02"), nil
}

func (date *Date) UnmarshalCSV(csv string) (err error) {
	date.Time, err = time.Parse("2006-01-02", csv)
	if err != nil {
		return err
	}
	return nil
}

type TimeEntry struct {
	Id        string  `csv:"id"`
	ProjectId string  `csv:"project_id"`
	TaskId    string  `csv:"task_id"`
	Hours     float64 `csv:"hours"`
	Notes     string  `csv:"notes"`
	Date      Date    `csv:"date"`
}

func main() {
	entriesFile, err := os.OpenFile("entries.csv", os.O_RDWR, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer entriesFile.Close()

	entries := []*TimeEntry{}

	if err := gocsv.UnmarshalFile(entriesFile, &entries); err != nil {
		panic(err)
	}

	for _, entry := range entries {
		request := struct {
			XMLName   xml.Name             `xml:"request"`
			Method    string               `xml:"method,attr"`
			TimeEntry freshbooks.TimeEntry `xml:"time_entry"`
		}{
			Method: "time_entry.create",
			TimeEntry: freshbooks.TimeEntry{
				ProjectId: entry.ProjectId,
				TaskId:    entry.TaskId,
				Hours:     entry.Hours,
				Notes:     entry.Notes,
				Date:      freshbooks.Date{entry.Date.Time},
			},
		}
		response, err := freshbooks.Do(request)
		fmt.Printf("%+v\n%+s\n", err, response)
	}
}
