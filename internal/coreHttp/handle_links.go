package corehttp

import (
	"hish22/grpm/internal/config"
	"net/url"
	"path/filepath"
	"strings"

	charmlog "github.com/charmbracelet/log"
)

const (
	BaseLink                   = "https://github.com/"
	ApiLink                    = "https://api.github.com/"
	BaseSearchLink             = "https://github.com/search?q="
	RepoQuery                  = "&type=repositories"
	PageQuery                  = "&p="
	MostStarsQuery             = "&s=stars&o=desc"
	FewStarsQuery              = "&s=stars&o=asc"
	ReleasesEndPoint           = "releases"
	ReposEndPoint              = "repos"
	ReposLatestReleaseEndPoint = "latest"
	ReposTagsEndPoint          = "tags"
	LatestFiveReleasesQuery    = "?per_page=5"
)

type RequestLink struct {
	Base      string
	Endpoints []string
	Queries   []string
}

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

func WriteDownloadsFilePath(filename string) string {
	downlaodsPath, err := config.GrpmDownloadedDirPath()
	if err != nil {
		charmlog.Error("Failed to return download path", "error", err)
	}
	return filepath.Join(downlaodsPath.String(), filename)
}

func WriteLibFilePath(filename string) string {
	libPath, err := config.GrpmLibraryDirPath()
	if err != nil {
		charmlog.Error("Failed to return library path", "error", err)
	}
	return filepath.Join(libPath.String(), filename)
}
