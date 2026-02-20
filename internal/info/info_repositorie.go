package info

import (
	corehttp "hish22/grpm/internal/coreHttp"
	"hish22/grpm/internal/persistance"
	"hish22/grpm/internal/structures"

	charmlog "github.com/charmbracelet/log"
)

func InfoRepository(owner string, name string) *structures.Repository {
	var repository structures.Repository
	link := corehttp.RequestLink{
		Base:      corehttp.ApiLink,
		Endpoints: []string{"repos", owner, name},
	}.Build()
	if !persistance.FetchFromCache(&repository, link) {
		if err := corehttp.Request(link, &repository); err != nil {
			charmlog.Error("Failed to info specified repository", "error", err)
			return &structures.Repository{}
		}
		persistance.NewCache(link, &repository)
	}
	return &repository
}
