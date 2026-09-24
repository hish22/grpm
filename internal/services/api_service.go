package services

/*
 * API service.
 *
 * Provide a general api to request
 * different version control platforms.
 */

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

// General structure for each api request.
type ApiService struct {
	Link    *LinkService
	Headers map[string]string
	Method  string
	Timeout time.Duration
}

// Create a new api request.
func CreateApiRequest(link *LinkService, headers map[string]string, method string, timeout time.Duration) (*ApiService, error) {
	if link.IsEmpty() {
		return nil, errors.New("Failed to create new Api Request, LinkService object must not be empty")
	}
	return &ApiService{
		link,
		headers,
		method,
		timeout,
	}, nil
}

// Perform a new http request to an api platform.
func (api ApiService) Request(ctx context.Context) (*http.Response, error) {
	// Build a link string instance.
	url, err := api.Link.String()

	if err != nil {
		return nil, err
	}

	// Request an api with context.
	req, err := http.NewRequestWithContext(ctx, api.Method, url, nil)
	if err != nil {
		return nil, err
	}

	// Register provided headers.
	for key, header := range api.Headers {
		req.Header.Set(key, header)
	}

	// Perform the request and return the response.
	c := http.Client{}
	res, err := c.Do(req)

	// Through error when status code is 4xx or 5xx.
	if res.StatusCode >= 400 || res.StatusCode >= 500 {
		return nil, errors.New("Failed to perform request, " + res.Status)
	}

	return res, err
}

// Perform a new http request to an api platform, with handling decoding and context operations.
func (api ApiService) RequestWithDecode(structure any) error {
	// Create a context with a timeout duration.
	ctx, cancel := context.WithTimeout(context.Background(), api.Timeout)
	defer cancel()

	// Perform new request to api platform.
	res, err := api.Request(ctx)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	// Perform json decode into provided structure.
	err = json.NewDecoder(res.Body).Decode(&structure)
	if err != nil {
		return err
	}
	return nil
}
