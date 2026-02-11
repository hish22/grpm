package corehttp

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	charmlog "github.com/charmbracelet/log"
)

func Request(link string, structure any) {
	response, err := newHttpRequest(link)
	if err != nil {
		charmlog.Fatal("Failed to request", "link", link)
	}
	defer response.Body.Close()
	err = json.NewDecoder(response.Body).Decode(&structure)
	if err != nil {
		charmlog.Fatal("Failed to read response body of", "link", link)
	}
}

func newHttpRequest(link string) (*http.Response, error) {
	req, _ := http.NewRequest("GET", link, nil)
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("User-Agent", "grpm/0.0.1")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Can't do an http request, ", err)
	}
	if resp.StatusCode == 200 {
		return resp, nil
	}
	return nil, errors.New("Error with status code: " + strconv.Itoa(resp.StatusCode))
}

func NewJsonReq(link string) (*http.Response, error) {
	req, _ := http.NewRequest("GET", link, nil)
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("User-Agent", "grpm/0.0.1")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Can't do an http request, ", err)
	}
	if resp.StatusCode == 200 {
		return resp, nil
	}
	return nil, errors.New("Error with status code: " + strconv.Itoa(resp.StatusCode))
}
