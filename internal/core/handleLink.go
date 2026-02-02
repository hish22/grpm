package core

import (
	"fmt"
	"log"
	"net/url"
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

func InfoLink(owner *string, repo *string) string {
	url, err := url.JoinPath(BaseLink, *owner, *repo)
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
	fmt.Println(url)
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
