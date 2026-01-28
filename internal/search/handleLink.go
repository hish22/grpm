package search

import (
	"hish22/grpm/internal/packet"
)

const (
	SearchLink     = "https://github.com/search?q="
	RepoQuery      = "&type=repositories"
	PageQuery      = "&p="
	MostStarsQuery = "&s=stars&o=desc"
	FewStarsQuery  = "&s=stars&o=asc"
)

func searchLink(repo *packet.RepoInfo) string {
	var link string
	if repo.MostStars {
		link = SearchLink + repo.Name + RepoQuery + PageQuery + repo.Page + MostStarsQuery
	} else if repo.FewStars {
		link = SearchLink + repo.Name + RepoQuery + PageQuery + repo.Page + FewStarsQuery
	} else {
		link = SearchLink + repo.Name + RepoQuery + PageQuery + repo.Page
	}

	return link

}
