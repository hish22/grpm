package setup

import (
	"compress/gzip"
	"os"

	charmlog "github.com/charmbracelet/log"
)

func unzipFileTarGz(repo string, location string, assetID int) {
	// open the compressed file
	file, err := os.Open(location)
	if err != nil {
		charmlog.Error("Failed to open compressed file tar.gz", "error", err)
		return
	}
	defer file.Close()
	// Uncompress the file of .gz
	gzip, err := gzip.NewReader(file)
	if err != nil {
		charmlog.Error("Failed to uncompress .gz file", "error", err)
		return
	}
	defer gzip.Close()
	tarReader(repo, gzip, location, assetID)
}
