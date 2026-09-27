package gitlab

import (
	"fmt"
	"hish22/grpm/internal/services"
	"testing"
	"time"
)

func TestRequestOfAGitlabProjectReleases(t *testing.T) {
	link, err := services.CreateRequestLink(
		services.GitlabApiBase,
		[]string{"projects", "10174980", "releases"},
		[]string{},
	)

	fmt.Println(link.String())

	if err != nil {
		t.Error("Failed to create a reqest link")
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
		t.Error("Failed to create a new api request to get a gitlab release")
		return
	}

	// Gitlab Project instance.
	r := []release{}

	err = apiReq.RequestWithDecode(&r)

	if err != nil {
		t.Error("Failed to request the gitlab release page with decode")
		return
	}

	// Print project.
	fmt.Println(r)
}
