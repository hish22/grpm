package services

/*
 * Link Service
 *
 * Provide the ability to construct
 * manage and build a string for a
 * request api link.
 *
 */

import (
	"net/url"
	"strings"
)

const (
	GithubApiBase = "https://api.github.com/"
)

type RequestLink struct {
	Base      string
	Endpoints []string
	Queries   []string
}

// Create a new request link and return the pointer to it.
func CreateRequestLink(b string, ep []string, q []string) *RequestLink {
	return &RequestLink{
		b,
		ep,
		q,
	}
}

// Build a link from RequestLink object into a string.
func (link RequestLink) String() (string, error) {
	// Construct the link without the queries.
	joinedLink, err := url.JoinPath(link.Base, link.Endpoints...)
	if err != nil {
		return "", err
	}
	// Add queries only if queries len is greater than 0.
	if len(link.Queries) > 0 {
		queryPack := "?" + strings.Join(link.Queries, "&")
		joinedLink = joinedLink + queryPack
		return joinedLink, nil
	}
	return joinedLink, nil
}
