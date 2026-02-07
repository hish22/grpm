package structures

type Assets struct {
	ID          int    `json:"id"`
	Url         string `json:"url"`
	AssetName   string `json:"name"`
	Size        int    `json:"size"`
	Digest      string `json:"digest"`
	ContentType string `json:"content_type"`
	DownloadUrl string `json:"browser_download_url"`
}

type TrackedAsset struct {
	ID          int
	RepoName    string
	AssetName   string
	Location    string
	Tag         string
	ReleaseName string
	Size        int
	Digest      string
}
