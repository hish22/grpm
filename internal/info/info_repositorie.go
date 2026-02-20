package info

import (
	"fmt"
	corehttp "hish22/grpm/internal/coreHttp"
	"hish22/grpm/internal/persistance"
	"hish22/grpm/internal/structures"
)

func InfoRepository(owner string, name string) (*structures.Repository, error) {
	var repository structures.Repository
	link := corehttp.RequestLink{
		Base:      corehttp.ApiLink,
		Endpoints: []string{"repos", owner, name},
	}.Build()
	if !persistance.FetchFromCache(&repository, link) {
		if err := corehttp.Request(link, &repository); err != nil {
			return &structures.Repository{}, fmt.Errorf("Failed to search specified repository")
		}
		persistance.NewCache(link, &repository)
	}
	return &repository, nil
}
