package setup

import (
	"compress/gzip"
	"os"

	charmlog "github.com/charmbracelet/log"
)

func unzipFileTarGz(location string, assetID int) {
	// open the compressed file
	file, err := os.Open(location)
	if err != nil {
		charmlog.Fatal("Failed to open compressed file tar.gz", "error", err)
	}
	defer file.Close()
	// Uncompress the file of .gz
	gzip, err := gzip.NewReader(file)
	if err != nil {
		charmlog.Fatal("Failed to uncompress .gz file", "error", err)
	}
	defer gzip.Close()
	tarReader(gzip, location, assetID)
}
