package github

import (
	"fmt"
	"hish22/grpm/internal/services"
	"testing"
	"time"
)

func TestRequestOfAGithubRelease(t *testing.T) {
	link, err := services.CreateRequestLink(
		services.GithubApiBase,
		[]string{"repos", "hish22", "depoty", "releases"},
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
		t.Error("Failed to create a new api request to get a github repo", err)
		return
	}

	// Gitlab Project instance.
	re := []release{}

	err = apiReq.RequestWithDecode(&re)

	if err != nil {
		t.Error("Failed to request the github repo page with decode", err)
		return
	}

	// Print project.
	fmt.Println(re)
}

func TestRequestRepoReleasesByOwnerName(t *testing.T) {
	rs := []release{}
	if err := RequestRepoReleasesByOwnerName(&rs, "hish22", "depoty"); err != nil {
		fmt.Println(err)
	}
	fmt.Println(rs)
}
