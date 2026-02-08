package util

import (
	"hish22/grpm/internal/structures"
	"testing"
)

func TestArchitectureAssetsMatch_AMD64(t *testing.T) {
	amd64Arch := "amd64"
	matchedReleases := make([]structures.Assets, 0)

	assets := []structures.Assets{
		{AssetName: "myapp-linux-amd64.tar.gz"},
		{AssetName: "myapp-linux-x86-64.deb"},
		{AssetName: "myapp-linux-arm64.tar.gz"},
		{AssetName: "myapp-windows-x64.exe"},
		{AssetName: "myapp-darwin-Intel64.zip"},
	}

	for i := range assets {
		ArchitectureAssetsMatch(&amd64Arch, &assets[i], &matchedReleases)
	}

	if len(matchedReleases) != 4 {
		t.Errorf("Expected 4 matches for amd64, got %d", len(matchedReleases))
	}

	expectedNames := map[string]bool{
		"myapp-linux-amd64.tar.gz": true,
		"myapp-linux-x86-64.deb":   true,
		"myapp-windows-x64.exe":    true,
		"myapp-darwin-Intel64.zip": true,
	}

	for _, asset := range matchedReleases {
		if !expectedNames[asset.AssetName] {
			t.Errorf("Unexpected asset name: %s", asset.AssetName)
		}
	}
}

func TestArchitectureAssetsMatch_ARM64(t *testing.T) {
	arm64Arch := "arm64"
	matchedReleases := make([]structures.Assets, 0)

	assets := []structures.Assets{
		{AssetName: "myapp-linux-aarch64.tar.gz"},
		{AssetName: "myapp-linux-armv8-a.deb"},
		{AssetName: "myapp-linux-armv9-a.tar.gz"},
		{AssetName: "myapp-linux-amd64.tar.gz"},
		{AssetName: "myapp-darwin-arm64.zip"},
	}

	for i := range assets {
		ArchitectureAssetsMatch(&arm64Arch, &assets[i], &matchedReleases)
	}

	if len(matchedReleases) != 3 {
		t.Errorf("Expected 3 matches for arm64, got %d", len(matchedReleases))
	}

	expectedNames := map[string]bool{
		"myapp-linux-aarch64.tar.gz": true,
		"myapp-linux-armv8-a.deb":    true,
		"myapp-linux-armv9-a.tar.gz": true,
	}

	for _, asset := range matchedReleases {
		if !expectedNames[asset.AssetName] {
			t.Errorf("Unexpected asset name: %s", asset.AssetName)
		}
	}
}

func TestArchitectureAssetsMatch_386(t *testing.T) {
	i386Arch := "386"
	matchedReleases := make([]structures.Assets, 0)

	assets := []structures.Assets{
		{AssetName: "myapp-linux-i386.deb"},
		{AssetName: "myapp-windows-amd64.exe"},
		{AssetName: "myapp-linux-arm64.tar.gz"},
	}

	for i := range assets {
		ArchitectureAssetsMatch(&i386Arch, &assets[i], &matchedReleases)
	}

	if len(matchedReleases) != 1 {
		t.Errorf("Expected 1 match for 386, got %d", len(matchedReleases))
	}

	if matchedReleases[0].AssetName != "myapp-linux-i386.deb" {
		t.Errorf("Expected myapp-linux-i386.deb, got %s", matchedReleases[0].AssetName)
	}
}

func TestArchitectureAssetsMatch_UnknownArch(t *testing.T) {
	unknownArch := "mips"
	matchedReleases := make([]structures.Assets, 0)

	assets := []structures.Assets{
		{AssetName: "myapp-linux-amd64.tar.gz"},
		{AssetName: "myapp-linux-arm64.tar.gz"},
	}

	for i := range assets {
		ArchitectureAssetsMatch(&unknownArch, &assets[i], &matchedReleases)
	}

	if len(matchedReleases) != 0 {
		t.Errorf("Expected 0 matches for unknown architecture, got %d", len(matchedReleases))
	}
}

func TestArchitectureAssetsMatch_NoMatch(t *testing.T) {
	amd64Arch := "amd64"
	matchedReleases := make([]structures.Assets, 0)

	assets := []structures.Assets{
		{AssetName: "myapp-linux-arm64.tar.gz"},
		{AssetName: "myapp-darwin-arm64.zip"},
	}

	for i := range assets {
		ArchitectureAssetsMatch(&amd64Arch, &assets[i], &matchedReleases)
	}

	if len(matchedReleases) != 0 {
		t.Errorf("Expected 0 matches when no architecture matches, got %d", len(matchedReleases))
	}
}
