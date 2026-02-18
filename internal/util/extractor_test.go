package util

import "testing"

func TestExtensionExtractor(t *testing.T) {
	file := "jolt-1.2.0-linux-amd64.tar.gz"
	extensions := ExtensionExtractor(file)

	if extensions[0] != "tar" {
		t.Errorf("Expected tar as first extension, got:%s", extensions[0])
	}

	if extensions[1] != "gz" {
		t.Errorf("Expected tar as second extension, got:%s", extensions[1])
	}

}

func TestNameAndExtensionExtractor(t *testing.T) {
	file := "fzf-0.67.0-windows_amd64.zip"
	result := NameAndExtensionExtractor(file)

	if result[0] != "fzf-0.67.0-windows_amd64.zip" {
		t.Errorf("Expected fzf-0.67.0-windows_amd64.zip as first extension, got:%s", result[0])
	}

	if result[1] != "fzf-0.67.0-windows_amd64" {
		t.Errorf("Expected fzf-0.67.0-windows_amd64 as first extension, got:%s", result[1])
	}
}
