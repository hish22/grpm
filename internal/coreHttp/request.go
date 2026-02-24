package corehttp

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	charmlog "github.com/charmbracelet/log"
)

type ApiRequest struct {
	Link    RequestLink
	Timeout time.Duration
}

func (instance ApiRequest) RequestWithDecode(structure any) error {
	BuiltLink := instance.Link.Build()
	ctx, cancel := context.WithTimeout(context.Background(), instance.Timeout)
	defer cancel()
	response, err := instance.httpRequest(ctx, BuiltLink)

	if err != nil {
		charmlog.Error("Failed to request", "link", BuiltLink)
		return err
	}
	defer response.Body.Close()
	err = json.NewDecoder(response.Body).Decode(&structure)
	if err != nil {
		charmlog.Error("Failed to read response body of", "link", BuiltLink, "error", err)
		return err
	}
	return nil
}

func (instance ApiRequest) RequestWithContext(ctx context.Context) (*http.Response, error) {
	req, _ := http.NewRequestWithContext(ctx, "GET", instance.Link.Build(), nil)
	req.Header.Set("User-Agent", "grpm/0.0.1")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

/* Requesting github api, and return a http response */
func (instance ApiRequest) httpRequest(ctx context.Context, link string) (*http.Response, error) {
	req, _ := http.NewRequestWithContext(ctx, "GET", link, nil)
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("User-Agent", "grpm/0.0.1")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		charmlog.Error("Failed to perform http request, ", "error", err)
		return nil, errors.New("Deadline exceeded")
	}
	if resp.StatusCode == 200 {
		return resp, nil
	}
	return nil, errors.New("Error with status code: " + strconv.Itoa(resp.StatusCode))
}
