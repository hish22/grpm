package setup

import (
	"os"

	charmlog "github.com/charmbracelet/log"
	"github.com/klauspost/compress/zstd"
)

func unzipFileTarZst(location string, assetID int) {
	file, err := os.Open(location)
	if err != nil {
		charmlog.Fatal("Failed to open compressed file tar.zst", "error", err)
	}
	defer file.Close()
	zstgz, err := zstd.NewReader(file)
	if err != nil {
		charmlog.Fatal("Failed to uncompress .zst file", "error", err)
	}
	tarReader(zstgz, location, assetID)
}
