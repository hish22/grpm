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
	Base      string
	Endpoints []string
	Queries   []string
}

// Create a new request link and return the pointer to it.
func CreateRequestLink(b string, ep []string, q []string) (*LinkService, error) {
	if len(b) == 0 {
		return nil, errors.New("Failed to create new RequestLink, Base must not be of length 0")
	}
	return &LinkService{
		b,
		ep,
		q,
	}, nil
}

// Provide a way to make sure that LinkService is empty object or not (zero values)
func (link LinkService) IsEmpty() bool {
	return len(link.Base) == 0 && len(link.Endpoints) == 0 && len(link.Queries) == 0
}

// Build a link from RequestLink object into a string.
func (link LinkService) String() (string, error) {
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
