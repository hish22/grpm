package info

import (
	corehttp "hish22/grpm/internal/coreHttp"
	"hish22/grpm/internal/persistance"
	"hish22/grpm/internal/structures"
)

func InfoRepository(owner *string, name *string) *structures.Repository {
	var repository structures.Repository
	link := corehttp.RequestLink{
		Base:      corehttp.ApiLink,
		Endpoints: []string{"repos", *owner, *name},
	}.Build()
	if !persistance.FetchFromCache(&repository, *link) {
		corehttp.Request(link, &repository)
		persistance.NewCache(*link, &repository)
	}
	return &repository
}
