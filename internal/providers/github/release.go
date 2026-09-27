package github

import (
	"hish22/grpm/internal/services"
	"time"
)

/*
 * Release Information struct
 *
 * A struct that provides a basic release
 * information for any repository contained
 * on the github platform.
 */

type release struct {
	Url         string    `json:"html_url"`
	ID          int       `json:"id"`
	TagName     string    `json:"tag_name"`
	Name        string    `json:"name"`
	Description string    `json:"body"`
	CreatedAt   time.Time `json:"created_at"`
	Author      owner     `json:"author"`
	Assets      []asset   `json:"assets"`
}

func RequestRepoReleasesByOwnerName(r *[]release, owner string, name string) error {
	link, err := services.CreateRequestLink(
		services.GithubApiBase,
		[]string{"repos", owner, name, "releases"},
		[]string{},
	)

	if err != nil {
		return err
	}

	apiReq, err := services.CreateApiRequest(
		link,
		map[string]string{
			"Accept":     "application/json",
			"User-Agent": "grpm/0.0.1",
		},
		"GET",
		5*time.Second,
	)

	if err != nil {
		return err
	}

	err = apiReq.RequestWithDecode(&r)

	if err != nil {
		return err
	}

	return nil
}
