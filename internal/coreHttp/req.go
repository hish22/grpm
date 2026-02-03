package corehttp

import (
	"errors"
	"log"
	"net/http"
	"strconv"
)

func NewJsonReq(link *string) (*http.Response, error) {
	req, _ := http.NewRequest("GET", *link, nil)
	req.Header.Set("Accept", "application/json")
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
