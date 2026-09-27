package github

/*
 * Asset Information struct
 *
 * A struct that provides a basic asset
 * information for any release contained
 * in any repository that on github platform.
 */

type asset struct {
	ID          int    `json:"id"`
	Url         string `json:"url"`
	DownloadUrl string `json:"browser_download_url"`
	Name        string `json:"name"`
	Size        int    `json:"size"`
	Digest      string `json:"digest"`
}
