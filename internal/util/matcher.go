package util

import (
	"hish22/grpm/internal/structures"
	"strings"
)

func ArchitectureAssetsMatch(arch *string, asset *structures.Assets, matchedReleases *[]structures.Assets) {
	var archs []string
	switch *arch {
	case "amd64":
		archs = []string{"amd64", "x86", "x86-64", "Intel 64", "x64", "EM64T", "IA-32e", "64bit"}
	case "arm64":
		archs = []string{"aarch64", "armv8-a", "armv9-a"}
	case "386":
		archs = []string{"i386", "32bit"}
	default:
		archs = []string{}
	}
	for _, filteredArch := range archs {
		if strings.Contains(asset.AssetName, filteredArch) {
			*matchedReleases = append(*matchedReleases, *asset)
		}
	}
}
