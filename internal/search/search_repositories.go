package search

import (
	"fmt"
	corehttp "hish22/grpm/internal/coreHttp"
	"hish22/grpm/internal/persistance"
	"hish22/grpm/internal/structures"
	"time"
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
	}
	if !persistance.FetchFromCache(&repositories, link.Build()) {
		request := corehttp.ApiRequest{
			Link:    link,
			Timeout: time.Second * 10,
		}
		if err := request.RequestWithDecode(&repositories); err != nil {
			return &structures.Repositories{}, fmt.Errorf("Failed to search specified repository")
		}
		persistance.NewCache(link.Build(), &repositories)
	}

	if repositories.TotalCount != 0 {
		return &repositories, nil
	} else {
		return &repositories, fmt.Errorf("Failed to find any specified repository")
	}
}
