package github

import (
	"hish22/grpm/internal/services"
	"time"
)

/*
 * repository Information struct
 *
 * A struct that provides basic a building
 * information for any repository contained
 * on the github platform.
 */

type repository struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	FullName    string    `json:"full_name"`
	Description string    `json:"description"`
	Owner       owner     `json:"owner"`
	Url         string    `json:"html_url"`
	CreatedAt   time.Time `json:"created_at"`
}

func RequestRepoByOwnerName(r *repository, owner string, name string) error {
	link, err := services.CreateRequestLink(
		services.GithubApiBase,
		[]string{"repos", owner, name},
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
