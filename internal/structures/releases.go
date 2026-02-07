package structures

import "time"

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
