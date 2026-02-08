package structures

import (
	"encoding/json"
	"testing"
)

func TestAssets_JSONMarshal(t *testing.T) {
	asset := Assets{
		ID:          12345,
		Url:         "https://api.github.com/repos/owner/repo/releases/assets/12345",
		AssetName:   "myapp-linux-amd64.tar.gz",
		Size:        1024000,
		Digest:      "sha256:abc123",
		ContentType: "application/gzip",
		DownloadUrl: "https://github.com/owner/repo/releases/download/v1.0/myapp-linux-amd64.tar.gz",
	}

	data, err := json.Marshal(asset)
	if err != nil {
		t.Fatalf("Failed to marshal Assets: %v", err)
	}

	if len(data) == 0 {
		t.Error("Marshaled data is empty")
	}

	expectedJSON := `{"id":12345,"url":"https://api.github.com/repos/owner/repo/releases/assets/12345","name":"myapp-linux-amd64.tar.gz","size":1024000,"digest":"sha256:abc123","content_type":"application/gzip","browser_download_url":"https://github.com/owner/repo/releases/download/v1.0/myapp-linux-amd64.tar.gz"}`

	var actual, expected map[string]any
	json.Unmarshal(data, &actual)
	json.Unmarshal([]byte(expectedJSON), &expected)

	for key, expVal := range expected {
		if actVal, ok := actual[key]; !ok || expVal != actVal {
			t.Errorf("JSON mismatch for key %s: expected %v, got %v", key, expVal, actVal)
		}
	}
}

func TestAssets_JSONUnmarshal(t *testing.T) {
	jsonData := `{
		"id": 12345,
		"url": "https://api.github.com/repos/owner/repo/releases/assets/12345",
		"name": "myapp-linux-amd64.tar.gz",
		"size": 1024000,
		"digest": "sha256:abc123",
		"content_type": "application/gzip",
		"browser_download_url": "https://github.com/owner/repo/releases/download/v1.0/myapp-linux-amd64.tar.gz"
	}`

	var asset Assets
	err := json.Unmarshal([]byte(jsonData), &asset)
	if err != nil {
		t.Fatalf("Failed to unmarshal Assets: %v", err)
	}

	if asset.ID != 12345 {
		t.Errorf("Expected ID 12345, got %d", asset.ID)
	}
	if asset.AssetName != "myapp-linux-amd64.tar.gz" {
		t.Errorf("Expected AssetName myapp-linux-amd64.tar.gz, got %s", asset.AssetName)
	}
	if asset.Size != 1024000 {
		t.Errorf("Expected Size 1024000, got %d", asset.Size)
	}
	if asset.Digest != "sha256:abc123" {
		t.Errorf("Expected Digest sha256:abc123, got %s", asset.Digest)
	}
}

func TestTrackedAsset_Fields(t *testing.T) {
	tracked := TrackedAsset{
		ID:          12345,
		RepoName:    "owner/repo",
		AssetName:   "myapp-linux-amd64.tar.gz",
		Location:    "/usr/local/bin",
		Tag:         "v1.0.0",
		ReleaseName: "Release v1.0.0",
		Size:        1024000,
		Digest:      "sha256:abc123",
	}

	if tracked.ID != 12345 {
		t.Errorf("Expected ID 12345, got %d", tracked.ID)
	}
	if tracked.RepoName != "owner/repo" {
		t.Errorf("Expected RepoName owner/repo, got %s", tracked.RepoName)
	}
	if tracked.Location != "/usr/local/bin" {
		t.Errorf("Expected Location /usr/local/bin, got %s", tracked.Location)
	}
}
