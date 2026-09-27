package github

import (
	"fmt"
	"hish22/grpm/internal/services"
	"testing"
	"time"
)

func TestRequestOfAGithubRepo(t *testing.T) {
	link, err := services.CreateRequestLink(
		services.GithubApiBase,
		[]string{"repos", "hish22", "depoty"},
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
	r := repository{}

	err = apiReq.RequestWithDecode(&r)

	if err != nil {
		t.Error("Failed to request the github repo page with decode")
		return
	}

	// Print project.
	fmt.Println(r)
}

func TestRequestRepoByOwnerName(t *testing.T) {
	repo := repository{}
	if err := RequestRepoByOwnerName(&repo, "hish22", "depoty"); err != nil {
		fmt.Println(err)
	}
	fmt.Println(repo)
}
