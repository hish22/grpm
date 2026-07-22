package services

import (
	"fmt"
	"testing"
)

func TestRequestLink(t *testing.T) {
	link, err := CreateRequestLink(
		GithubApiBase,
		[]string{"search", "repositories"},
		[]string{"q=fzf"},
	)

	if err != nil {
		t.Error(err)
		return
	}

	// fmt.Println(link.Base, link.Endpoints, link.Queries)

	stringLink, err := link.String()

	if err != nil {
		t.Error("Failed to build string link", "error", err)
		return
	}

	fmt.Println(stringLink)

}

func TestRequestLinkWithEmptyBase(t *testing.T) {
	link, err := CreateRequestLink("", []string{}, []string{})

	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println(link.String())

}
