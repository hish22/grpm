package release

import (
	"hish22/grpm/internal/core"
	"hish22/grpm/internal/serialization"
	"io"
	"log"
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

func fetchReleases(link *string, structure any) {
	resp, err := core.NewJsonReq(link)
	if err != nil {
		log.Fatal("Can't fetch releases, ", err)
	}
	defer resp.Body.Close()
	buf, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Can't read releases buffer data, ", err)
	}
	serialization.JsonUnserialization(buf, &structure)
}

func FetchLatestReleases(repo *string) []Release {
	link := core.ReleasesLatestLink(repo)
	var jsonReleasesResult []Release
	fetchReleases(&link, &jsonReleasesResult)
	return jsonReleasesResult
}

func FetchSpecificRelease(repo *string, tag *string) *Release {
	link := core.ReleasesByTagLink(repo, tag)
	var jsonReleaseResult Release
	fetchReleases(&link, &jsonReleaseResult)
	return &jsonReleaseResult
}
