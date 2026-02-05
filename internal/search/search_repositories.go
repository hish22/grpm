package search

import (
	corehttp "hish22/grpm/internal/coreHttp"
	"hish22/grpm/internal/persistance"
	"hish22/grpm/internal/structures"
)

type RepoInfo struct {
	Name  string
	Page  string
	Sort  string
	Order string
}

/* Search Repositories by requesting api.github */
func SearchRepositories(metadata *RepoInfo) *structures.Repositories {
	var repositories structures.Repositories
	link := corehttp.RequestLink{
		Base:      corehttp.ApiLink,
		Endpoints: []string{"search", "repositories"},
		Queries: []string{"q=" + metadata.Name, "sort=" +
			metadata.Sort, "order=" + metadata.Order,
			"page=" + metadata.Page},
	}.Build()
	if !persistance.FetchFromCache(&repositories, *link) {
		corehttp.Request(link, &repositories)
		persistance.NewCache(*link, &repositories)
	}
	return &repositories
}
