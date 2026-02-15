package corehttp

import (
	"hish22/grpm/internal/config"
	"log"
	"net/url"
	"os"
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

func WriteFilePath(fileName string) string {
	configs := config.DecodeTOMLConfig()
	homePath, err := os.UserHomeDir()

	if err != nil {
		charmlog.Fatal("Failed to return home dir path", "error", err)
	}

	return filepath.Join(homePath, configs.Downloaded, fileName)
}

func WriteDownloadsFilePath(filename string) string {
	downlaodsPath, err := config.GrpmDownloadedDirPath()
	if err != nil {
		charmlog.Error("Failed to return download path", "error", err)
	}
	return filepath.Join(downlaodsPath, filename)
}

func WriteLibFilePath(filename string) string {
	libPath, err := config.GrpmLibraryDirPath()
	if err != nil {
		charmlog.Error("Failed to return library path", "error", err)
	}
	return filepath.Join(libPath, filename)
}

func SearchLink(name *string, page *string, mostStars *bool, fewStars *bool) string {
	var link string
	if *mostStars {
		link = BaseSearchLink + *name + RepoQuery + PageQuery + *page + MostStarsQuery
	} else if *fewStars {
		link = BaseSearchLink + *name + RepoQuery + PageQuery + *page + FewStarsQuery
	} else {
		link = BaseSearchLink + *name + RepoQuery + PageQuery + *page
	}

	return link
}

func InfoLink(owner string, repo string) string {
	url, err := url.JoinPath(BaseLink, owner, repo)
	if err != nil {
		log.Fatal("Can't create such a link", err)
	}
	return url
}

func ReleasesLatestLink(repo *string) string {
	url, err := url.JoinPath(ApiLink, ReposEndPoint, *repo, ReleasesEndPoint)
	if err != nil {
		log.Fatal("Can't create such a link", err)
	}
	return url + LatestFiveReleasesQuery
}

func ReleaseLatestLink(repo *string) string {
	url, err := url.JoinPath(ApiLink, ReposEndPoint, *repo, ReleasesEndPoint, ReposLatestReleaseEndPoint)
	if err != nil {
		log.Fatal("Can't create such a link", err)
	}
	return url
}

// https://api.github.com/repos/owner/repo/releases/tags/v1.2.3
func ReleasesByTagLink(repo *string, tag *string) string {
	url, err := url.JoinPath(ApiLink, ReposEndPoint, *repo, ReleasesEndPoint, ReposTagsEndPoint, *tag)
	if err != nil {
		log.Fatal("Can't create such a link", err)
	}
	return url
}
