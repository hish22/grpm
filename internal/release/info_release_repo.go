package release

import (
	corehttp "hish22/grpm/internal/coreHttp"
	"time"
)

type Assets struct {
	ID          int    `json:"id"`
	Url         string `json:"url"`
	AssetName   string `json:"name"`
	Size        int    `json:"size"`
	Digest      string `json:"digest"`
	ContentType string `json:"content_type"`
	DownloadUrl string `json:"browser_download_url"`
}

type Release struct {
	ID          int       `json:"id"`
	Url         string    `json:"url"`
	AssetsUrl   string    `json:"assets_url"`
	TagName     string    `json:"tag_name"`
	ReleaseName string    `json:"name"`
	Assets      []Assets  `json:"assets"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	HtmlUrl     string    `json:"html_url"`
}

func FetchLatestReleases(name *string) []Release {
	link := corehttp.RequestLink{
		Base:      corehttp.ApiLink,
		Endpoints: []string{"repos", *name, "releases"},
	}.Build()
	var releasesResult []Release
	corehttp.Request(link, &releasesResult)
	return releasesResult
}

func FetchLatestRelease(name *string) *Release {
	link := corehttp.RequestLink{
		Base:      corehttp.ApiLink,
		Endpoints: []string{"repos", *name, "releases", "latest"},
	}.Build()
	var releaseResult Release
	corehttp.Request(link, &releaseResult)
	return &releaseResult
}

func FetchSpecificRelease(name *string, tag *string) *Release {
	link := corehttp.RequestLink{
		Base:      corehttp.ApiLink,
		Endpoints: []string{"repos", *name, "releases", "tags", *tag},
	}.Build()
	var releaseResult Release
	corehttp.Request(link, &releaseResult)
	return &releaseResult
}
