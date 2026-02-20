package search

import (
	corehttp "hish22/grpm/internal/coreHttp"
	"hish22/grpm/internal/persistance"
	"hish22/grpm/internal/structures"

	charmlog "github.com/charmbracelet/log"
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
	if !persistance.FetchFromCache(&repositories, link) {
		if err := corehttp.Request(link, &repositories); err != nil {
			charmlog.Error("Failed to search specified repository", "error", err)
			return &structures.Repositories{}
		}
		persistance.NewCache(link, &repositories)
	}
	return &repositories
}
