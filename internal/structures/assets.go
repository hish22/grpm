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
	ID                int    `sql:"id"`
	RepoName          string `sql:"repo"`
	AssetName         string `sql:"asset_name"`
	Location          string `sql:"location"`
	Tag               string `sql:"tag"`
	ReleaseName       string `sql:"release_name"`
	Size              int    `sql:"size"`
	Digest            string `sql:"Digest"`
	SetupStatus       bool   `sql:"setup_track"`
	SymlinkName       string `sql:"symlink_orenv_name"`
	FileSetupLocation string `sql:"file_setup_location"`
}
