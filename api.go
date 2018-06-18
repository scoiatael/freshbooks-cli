package freshbooks

import (
	"bytes"
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"os"
)

type Request struct {
	XMLName xml.Name `xml:"request"`
	Method  string   `xml:"method,attr"`
}

func Do(request interface{}) ([]byte, error) {
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
