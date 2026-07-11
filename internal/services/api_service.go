package services

/*
 * API service
 *
 * Provide a general api to request
 * different version control platforms.
 */

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
)

// General structure for each api request.
type ApiRequest struct {
	Link    RequestLink
	Method  string
	Headers map[string]string
	Timeout time.Duration
}

// Create a new api request.
func CreateApiRequest(link RequestLink, headers map[string]string, method string, timeout time.Duration) *ApiRequest {
	return &ApiRequest{
		link,
		method,
		headers,
		timeout,
	}
}

// Perform a new http request to an api platform.
func (api ApiRequest) NewRequest(ctx context.Context) (*http.Response, error) {
	// Build a link string instance.
	url, err := api.Link.String()

	if err != nil {
		return nil, err
	}

	// request an api with context.
	req, err := http.NewRequestWithContext(ctx, api.Method, url, nil)

	if err != nil {
		return nil, err
	}

	// register provided headers.
	for key, header := range api.Headers {
		req.Header.Set(key, header)
	}

	// Perform the request and return the response.
	c := http.Client{}
	res, err := c.Do(req)
	return res, err
}

// Perform a new http request to an api platform, with handling decoding and context operations.
func (api ApiRequest) NewRequestWithDecode(structure any) error {
	// Create a context with a timeout duration.
	ctx, cancel := context.WithTimeout(context.Background(), api.Timeout)
	defer cancel()

	// Perform new request to api platform.
	res, err := api.NewRequest(ctx)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	// Perfom json decode into provided structure.
	err = json.NewDecoder(res.Body).Decode(&structure)
	if err != nil {
		return err
	}
	return nil
}
