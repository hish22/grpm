package gitlab

import (
	"hish22/grpm/internal/services"
	"time"
)

/*
 * Release Information struct
 *
 * A struct that provides a basic release
 * information for any project contained
 * on the gitlab platform.
 */

type release struct {
	TagName     string    `json:"tag_name"`
	Description string    `json:"description"`
	Name        string    `json:"name"`
	CreatedAt   time.Time `json:"created_at"`
	Released_at time.Time `json:"released_at"`
	Author      user      `json:"author"`
	Assets      asset     `json:"assets"`
}

func RequestProjectReleaseById(r *release, id string) error {
	link, err := services.CreateRequestLink(
		services.GitlabApiBase,
		[]string{"projects", id, "releases"},
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
