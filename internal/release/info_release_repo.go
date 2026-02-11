package release

import (
	corehttp "hish22/grpm/internal/coreHttp"
	"hish22/grpm/internal/structures"
)

func FetchLatestReleases(name *string) []structures.Release {
	link := corehttp.RequestLink{
		Base:      corehttp.ApiLink,
		Endpoints: []string{"repos", *name, "releases"},
	}.Build()
	var releasesResult []structures.Release
	corehttp.Request(link, &releasesResult)
	return releasesResult
}

func FetchLatestRelease(name string) *structures.Release {
	link := corehttp.RequestLink{
		Base:      corehttp.ApiLink,
		Endpoints: []string{"repos", name, "releases", "latest"},
	}.Build()
	var releaseResult structures.Release
	corehttp.Request(link, &releaseResult)
	return &releaseResult
}

func FetchSpecificRelease(name string, tag string) *structures.Release {
	link := corehttp.RequestLink{
		Base:      corehttp.ApiLink,
		Endpoints: []string{"repos", name, "releases", "tags", tag},
	}.Build()
	var releaseResult structures.Release
	corehttp.Request(link, &releaseResult)
	return &releaseResult
}
