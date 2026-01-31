package core

import (
	"errors"
	"log"
	"net/http"
	"strconv"
)

func NewJsonReq(link *string) (*http.Response, error) {
	req, _ := http.NewRequest("GET", *link, nil)
	req.Header.Set("Accept", "application/json")
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
