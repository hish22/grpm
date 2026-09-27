package gitlab

import (
	"fmt"
	"hish22/grpm/internal/services"
	"testing"
	"time"
)

func TestGetAUserFromGitlab(t *testing.T) {
	link, err := services.CreateRequestLink(
		services.GitlabApiBase,
		[]string{"users"},
		[]string{"username=paljort"},
	)

	fmt.Println(link.String())

	if err != nil {
		t.Error("Failed to create a reqest link", err)
		return
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
		t.Error("Failed to create a new api request to get a gitlab user", err)
		return
	}

	// Gitlab Project instance.
	u := []user{}

	err = apiReq.RequestWithDecode(&u)

	if err != nil {
		t.Error("Failed to request the gitlab user page with decode", err)
		return
	}

	// Print project.
	fmt.Println(u[0])
}

func TestRequestAUserFromGitlab(t *testing.T) {
	u := []user{}
	if err := RequestUserByUsername(&u, "jack_smith"); err != nil {
		fmt.Println(err)
	}
	fmt.Println(u)
}
