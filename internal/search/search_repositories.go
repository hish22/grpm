package search

import (
	"fmt"
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
func SearchRepositories(metadata *RepoInfo) (*structures.Repositories, error) {
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
			return &structures.Repositories{}, fmt.Errorf("Failed to search specified repository")
		}
		persistance.NewCache(link, &repositories)
	}

	if repositories.TotalCount != 0 {
		return &repositories, nil
	} else {
		return &repositories, fmt.Errorf("Couldn't find any specified repository")
	}
}
