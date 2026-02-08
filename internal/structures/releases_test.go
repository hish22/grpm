package structures

import (
	"encoding/json"
	"testing"
	"time"
)

func TestRelease_JSONMarshal(t *testing.T) {
	releaseTime := time.Date(2024, 1, 15, 10, 30, 0, 0, time.UTC)
	release := Release{
		ID:          12345,
		Url:         "https://api.github.com/repos/owner/repo/releases/12345",
		AssetsUrl:   "https://api.github.com/repos/owner/repo/releases/12345/assets",
		TagName:     "v1.0.0",
		ReleaseName: "Release v1.0.0",
		Assets: []Assets{
			{
				ID:          111,
				AssetName:   "myapp-linux-amd64.tar.gz",
				Size:        1024000,
				DownloadUrl: "https://github.com/owner/repo/releases/download/v1.0.0/myapp-linux-amd64.tar.gz",
			},
			{
				ID:          222,
				AssetName:   "myapp-linux-arm64.tar.gz",
				Size:        998000,
				DownloadUrl: "https://github.com/owner/repo/releases/download/v1.0.0/myapp-linux-arm64.tar.gz",
			},
		},
		CreatedAt: releaseTime,
		UpdatedAt: releaseTime,
		HtmlUrl:   "https://github.com/owner/repo/releases/tag/v1.0.0",
	}

	data, err := json.Marshal(release)
	if err != nil {
		t.Fatalf("Failed to marshal Release: %v", err)
	}

	if len(data) == 0 {
		t.Error("Marshaled data is empty")
	}

	var result Release
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatalf("Failed to unmarshal Release: %v", err)
	}

	if result.TagName != "v1.0.0" {
		t.Errorf("Expected TagName v1.0.0, got %s", result.TagName)
	}
	if len(result.Assets) != 2 {
		t.Errorf("Expected 2 assets, got %d", len(result.Assets))
	}
}

func TestRelease_JSONUnmarshal(t *testing.T) {
	jsonData := `{
		"id": 12345,
		"url": "https://api.github.com/repos/owner/repo/releases/12345",
		"assets_url": "https://api.github.com/repos/owner/repo/releases/12345/assets",
		"tag_name": "v2.0.0",
		"name": "Release v2.0.0",
		"assets": [
			{
				"id": 111,
				"name": "myapp-linux-amd64.tar.gz",
				"size": 1024000
			}
		],
		"created_at": "2024-01-15T10:30:00Z",
		"updated_at": "2024-01-15T10:30:00Z",
		"html_url": "https://github.com/owner/repo/releases/tag/v2.0.0"
	}`

	var release Release
	err := json.Unmarshal([]byte(jsonData), &release)
	if err != nil {
		t.Fatalf("Failed to unmarshal Release: %v", err)
	}

	if release.ID != 12345 {
		t.Errorf("Expected ID 12345, got %d", release.ID)
	}
	if release.TagName != "v2.0.0" {
		t.Errorf("Expected TagName v2.0.0, got %s", release.TagName)
	}
	if len(release.Assets) != 1 {
		t.Errorf("Expected 1 asset, got %d", len(release.Assets))
	}
	if release.Assets[0].AssetName != "myapp-linux-amd64.tar.gz" {
		t.Errorf("Expected asset name myapp-linux-amd64.tar.gz, got %s", release.Assets[0].AssetName)
	}
}

func TestRelease_EmptyAssets(t *testing.T) {
	release := Release{
		ID:        12345,
		TagName:   "v1.0.0",
		Assets:    []Assets{},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	data, err := json.Marshal(release)
	if err != nil {
		t.Fatalf("Failed to marshal Release with empty assets: %v", err)
	}

	var result Release
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatalf("Failed to unmarshal Release with empty assets: %v", err)
	}

	if len(result.Assets) != 0 {
		t.Errorf("Expected 0 assets, got %d", len(result.Assets))
	}
}
