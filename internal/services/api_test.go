package services

import (
	"fmt"
	"testing"
	"time"
)

type Owner struct {
	ID       int    `json:"id"`
	Username string `json:"login"`
	Url      string `json:"url"`
	HtmlUrl  string `json:"html_url"`
	Type     string `json:"type"`
}

type license struct {
	Key    string `json:"key"`
	Name   string `json:"name"`
	SpdxID string `json:"spdx_id"`
	Url    string `json:"url"`
}

type Repository struct {
	ID                  int      `json:"id"`
	NodeID              string   `json:"node_id"`
	Name                string   `json:"name"`
	FullName            string   `json:"full_name"`
	Private             bool     `json:"private"`
	Owner               Owner    `json:"owner"`
	HtmlUrl             string   `json:"html_url"`
	Description         string   `json:"description"`
	Stars               int      `json:"stargazers_count"`
	Watchers            int      `json:"watchers_count"`
	Forks               int      `json:"forks"`
	ProgrammingLanguage string   `json:"language"`
	License             license  `json:"license"`
	Topics              []string `json:"topics"`
}

type Repositories struct {
	TotalCount   int          `json:"total_count"`
	Repositories []Repository `json:"items"`
}

func TestApiServiceWithGithubSearch(t *testing.T) {
	link, err := CreateRequestLink(
		GithubApiBase,
		[]string{"search", "repositories"},
		[]string{"q=fzf"},
	)

	if err != nil {
		t.Error(err)
		return
	}

	apiReq, err := CreateApiRequest(
		link,
		map[string]string{
			"Accept":     "application/vnd.github+json",
			"User-Agent": "grpm/0.0.1",
		},
		"GET",
		5*time.Second,
	)

	if err != nil {
		t.Error(err)
		return
	}

	var repos Repositories

	err = apiReq.RequestWithDecode(&repos)

	if err != nil {
		t.Error("Failed to fetch response and decode searched repo into struct", "error", err)
		return
	}

	fmt.Println(repos.TotalCount)
	fmt.Println(repos.Repositories)

}

func TestApiServiceWithGithubFailSearch(t *testing.T) {
	link, err := CreateRequestLink(
		GithubApiBase,
		[]string{"search", "repositories"},
		[]string{"q=awdawdawdawdawgfwgerg"},
	)

	if err != nil {
		t.Error(err)
		return
	}

	apiReq, err := CreateApiRequest(
		link,
		map[string]string{
			"Accept":     "application/vnd.github+json",
			"User-Agent": "grpm/0.0.1",
		},
		"GET",
		5*time.Second,
	)

	if err != nil {
		t.Error(err)
		return
	}

	var repos Repositories

	err = apiReq.RequestWithDecode(&repos)

	if err != nil {
		t.Error("Failed to fetch response and decode searched repo into struct", "error", err)
		return
	}

	if repos.TotalCount <= 0 {
		t.Error("No repos have been found")
		return
	}

	fmt.Println(repos.TotalCount)
	fmt.Println(repos.Repositories)

}
