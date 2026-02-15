package release

import (
	corehttp "hish22/grpm/internal/coreHttp"
	"hish22/grpm/internal/structures"
)

func FetchLatestReleases(name *string) ([]structures.Release, error) {
	link := corehttp.RequestLink{
		Base:      corehttp.ApiLink,
		Endpoints: []string{"repos", *name, "releases"},
	}.Build()
	var releasesResult []structures.Release
	err := corehttp.Request(link, &releasesResult)
	if err != nil {
		return releasesResult, err
	}
	return releasesResult, nil
}

func FetchLatestRelease(name string) (*structures.Release, error) {
	link := corehttp.RequestLink{
		Base:      corehttp.ApiLink,
		Endpoints: []string{"repos", name, "releases", "latest"},
	}.Build()
	var releaseResult structures.Release
	err := corehttp.Request(link, &releaseResult)
	if err != nil {
		return &releaseResult, err
	}
	return &releaseResult, nil
}

func FetchSpecificRelease(name string, tag string) (*structures.Release, error) {
	link := corehttp.RequestLink{
		Base:      corehttp.ApiLink,
		Endpoints: []string{"repos", name, "releases", "tags", tag},
	}.Build()
	var releaseResult structures.Release
	err := corehttp.Request(link, &releaseResult)
	if err != nil {
		return &releaseResult, err
	}
	return &releaseResult, nil
}
