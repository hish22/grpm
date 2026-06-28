package release

import (
	corehttp "hish22/grpm/internal/coreHttp"
	"hish22/grpm/internal/structures"
	"time"
)

func FetchLatestReleases(name *string) ([]structures.Release, error) {
	link := corehttp.RequestLink{
		Base:      corehttp.ApiLink,
		Endpoints: []string{"repos", *name, "releases"},
	}
	var releasesResult []structures.Release
	request := corehttp.ApiRequest{
		Link:    link,
		Timeout: time.Second * 10,
	}
	err := request.RequestWithDecode(&releasesResult)
	if err != nil {
		return releasesResult, err
	}
	return releasesResult, nil
}

func FetchLatestRelease(name string) (*structures.Release, error) {
	link := corehttp.RequestLink{
		Base:      corehttp.ApiLink,
		Endpoints: []string{"repos", name, "releases", "latest"},
	}
	var releaseResult structures.Release
	request := corehttp.ApiRequest{
		Link:    link,
		Timeout: time.Second * 10,
	}
	err := request.RequestWithDecode(&releaseResult)
	if err != nil {
		return &releaseResult, err
	}
	return &releaseResult, nil
}

func FetchSpecificRelease(name string, tag string) (*structures.Release, error) {
	link := corehttp.RequestLink{
		Base:      corehttp.ApiLink,
		Endpoints: []string{"repos", name, "releases", "tags", tag},
	}
	var releaseResult structures.Release
	request := corehttp.ApiRequest{
		Link:    link,
		Timeout: time.Second * 10,
	}
	err := request.RequestWithDecode(&releaseResult)
	if err != nil {
		return &releaseResult, err
	}
	return &releaseResult, nil
}
