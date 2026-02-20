package util

import (
	"bytes"
	"hish22/grpm/internal/structures"
	"os"
	"strings"

	charmlog "github.com/charmbracelet/log"
)

func IsBinary(file string) bool {
	f, err := os.Open(file)
	if err != nil {
		charmlog.Error("Failed to open file", "file", f.Name(), "error", err)
	}
	defer f.Close()
	header := make([]byte, 4)
	_, err = f.Read(header)
	if err != nil {
		charmlog.Warn("Failed to read file header", "error", err)
	}

	elfMagic := []byte{0x7F, 'E', 'L', 'F'} // Magic number header of a file
	return bytes.Equal(header, elfMagic)
}

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
