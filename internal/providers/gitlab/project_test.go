package gitlab

import (
	"fmt"
	"hish22/grpm/internal/services"
	"testing"
	"time"
)

func TestRequestOfAGitlabProject(t *testing.T) {
	link, err := services.CreateRequestLink(
		services.GitlabApiBase,
		[]string{"projects", "10174980"},
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
		t.Error("Failed to create a new api request to get a gitlab project")
		return
	}

	// Gitlab Project instance.
	p := project{}

	err = apiReq.RequestWithDecode(&p)

	if err != nil {
		t.Error("Failed to request the gitlab project page with decode")
		return
	}

	// Print project.
	fmt.Println(p)
}

func TestRequestOfAGitlabProjectById(t *testing.T) {
	p := project{}
	RequestProjectById(&p, "10174980")
	fmt.Println(p)
}

// func TestRequestOfAGitlabProjectBySearch(t *testing.T) {
// 	p := []project{}
// 	RequestProjectBySearch(p, "Veloren", "Veloren")
// 	fmt.Println(p)
// }
