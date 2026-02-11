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
