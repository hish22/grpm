package gitlab

import (
	"fmt"
	"hish22/grpm/internal/services"
	"time"
)

/*
 * User Information struct
 *
 * A struct that provides a basic author
 * information for any user belong
 * to the gitlab platform.
 */

type user struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	State    string `json:"state"`
	Locked   bool   `json:"locked"`
	Avatar   string `json:"avatar_url"`
	WebUrl   string `json:"web_url"`
}

func RequestUserByUsername(u *[]user, username string) error {
	link, err := services.CreateRequestLink(
		services.GitlabApiBase,
		[]string{"users"},
		[]string{"username=" + username},
	)

	fmt.Println(link.String())

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

	err = apiReq.RequestWithDecode(&u)

	if err != nil {
		return err
	}

	return nil
}
