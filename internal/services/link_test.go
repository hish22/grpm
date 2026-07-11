package services

import (
	"fmt"
	"testing"
)

func TestRequestLink(t *testing.T) {
	link := CreateRequestLink(
		GithubApiBase,
		[]string{"search", "repositories"},
		[]string{"q=fzf"},
	)

	fmt.Println(link.Base, link.Endpoints, link.Queries)

	stringLink, err := link.String()

	if err != nil {
		t.Error("Failed to build string link", "error", err)
		return
	}

	fmt.Println(stringLink)

}
