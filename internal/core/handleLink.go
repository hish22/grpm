package core

import (
	"hish22/grpm/internal/packet"
	"log"
	"net/url"
)

const (
	BaseLink       = "https://github.com/"
	BaseSearchLink = "https://github.com/search?q="
	RepoQuery      = "&type=repositories"
	PageQuery      = "&p="
	MostStarsQuery = "&s=stars&o=desc"
	FewStarsQuery  = "&s=stars&o=asc"
)

func SearchLink(repo *packet.RepoInfo) string {
	var link string
	if repo.MostStars {
		link = BaseSearchLink + repo.Name + RepoQuery + PageQuery + repo.Page + MostStarsQuery
	} else if repo.FewStars {
		link = BaseSearchLink + repo.Name + RepoQuery + PageQuery + repo.Page + FewStarsQuery
	} else {
		link = BaseSearchLink + repo.Name + RepoQuery + PageQuery + repo.Page
	}

	return link
}

func InfoLink(owner *string, repo *string) string {
	url, err := url.JoinPath(BaseLink, *owner, *repo)
	if err != nil {
		log.Fatal("Can't create such link", err)
	}
	return url
}
