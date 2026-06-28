package corehttp

import (
	"net/url"
	"strings"

	charmlog "github.com/charmbracelet/log"
)

const (
	BaseLink = "https://github.com/"
	ApiLink  = "https://api.github.com/"
)

type RequestLink struct {
	Base      string
	Endpoints []string
	Queries   []string
}

// Construct a link as a string
func (link RequestLink) Build() string {
	// Construct the link without the queries.
	construct, err := url.JoinPath(link.Base, link.Endpoints...)
	if err != nil {
		charmlog.Fatal("Failed to construct a request URL link, ", "Error", err)
	}
	// Add queries only if queries len is greater than 0.
	if len(link.Queries) > 0 {
		queryPack := "?" + strings.Join(link.Queries, "&")
		httpLink := construct + queryPack
		return httpLink
	}
	return construct
}
