package info

import (
	"fmt"
	corehttp "hish22/grpm/internal/coreHttp"
	"hish22/grpm/internal/persistance"
	"hish22/grpm/internal/structures"
	"time"
)

func InfoRepository(owner string, name string) (*structures.Repository, error) {
	var repository structures.Repository
	link := corehttp.RequestLink{
		Base:      corehttp.ApiLink,
		Endpoints: []string{"repos", owner, name},
	}
	if !persistance.FetchFromCache(&repository, link.Build()) {
		request := corehttp.ApiRequest{
			Link:    link,
			Timeout: time.Second * 10,
		}
		if err := request.RequestWithDecode(&repository); err != nil {
			return &structures.Repository{}, fmt.Errorf("Failed to search specified repository")
		}
		persistance.NewCache(link.Build(), &repository)
	}
	return &repository, nil
}
