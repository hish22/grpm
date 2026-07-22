package services

/*
 * Link Service.
 *
 * Provide the ability to construct
 * manage and build a string for a
 * request api link.
 *
 */

import (
	"errors"
	"net/url"
	"strings"
)

const (
	GithubApiBase = "https://api.github.com/"
)

type LinkService struct {
	base      string
	endpoints []string
	queries   []string
}

// Create a new request link. RequestLink expects a base, list of endpoints and a list of queries.
func CreateRequestLink(base string, endpoints []string, queries []string) (*LinkService, error) {
	if len(base) == 0 {
		return nil, errors.New("Failed to create new RequestLink, Base must not be of length 0")
	}
	return &LinkService{
		base,
		endpoints,
		queries,
	}, nil
}

// Provide a way to make sure that LinkService is empty object or not (zero values)
func (l LinkService) IsEmpty() bool {
	return len(l.base) == 0 && len(l.endpoints) == 0 && len(l.queries) == 0
}

// Build a link from RequestLink object into a string.
func (l LinkService) String() (string, error) {
	// Construct the link without the queries.
	joinedLink, err := url.JoinPath(l.base, l.endpoints...)
	if err != nil {
		return "nil", err
	}
	// Add queries only if queries len is greater than 0.
	if len(l.queries) > 0 {
		queryPack := "?" + strings.Join(l.queries, "&")
		joinedLink = joinedLink + queryPack
		return joinedLink, nil
	}
	return joinedLink, nil
}
