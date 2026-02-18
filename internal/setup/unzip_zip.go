package setup

import (
	"archive/zip"
	"hish22/grpm/internal/asset"
	"hish22/grpm/internal/config"
	"hish22/grpm/internal/util"
	"io"
	"os"
	"path/filepath"

	charmlog "github.com/charmbracelet/log"
)

func unzipZip(location string, assetID int) {
	zipf, err := zip.OpenReader(location)
	if err != nil {
		charmlog.Error("Failed to open zip file", "error", err)
		return
	}
	defer zipf.Close()
	// Fetch grpm lib path
	libLink, err := config.GrpmLibraryDirPath()
	if err != nil {
		charmlog.Error("Failed to fetch grpm lib path", "error", err)
	}

	// open the .zip file
	zFileInfo, err := os.Open(location)
	if err != nil {
		charmlog.Error("Failed to open file", "error", err, "file", location)
	}
	// Extract the file name without the ext
	name := util.NameAndExtensionExtractor(filepath.Base(zFileInfo.Name()))[1]

	// create file name
	parentDir := filepath.Join(libLink.String(), name)

	// Create the parent directory
	err = os.MkdirAll(parentDir, 0755)
	if err != nil {
		charmlog.Error("Failed to create direcotry", "error", err)
	}

	// track setup location
	asset.InsertFileSetupLocation(parentDir, assetID)

	// Create files into parent directory
	for _, f := range zipf.File {
		charmlog.Info(f.Name)
		// Open the file
		rc, err := f.Open()
		if err != nil {
			charmlog.Error("Failed to open file on .zip", "error", err)
		}
		defer rc.Close()

		// create file name
		assetPath := filepath.Join(parentDir, f.Name)

		if f.FileInfo().IsDir() {
			err := os.MkdirAll(assetPath, 0755)
			if err != nil {
				charmlog.Error("Failed to create direcotry", "error", err)
			}
		} else {
			nf, err := os.Create(assetPath)
			if err != nil {
				charmlog.Error("Failed to create a file", "error", err, "file", nf.Name())
			}
			_, err = io.Copy(nf, rc)
			if err != nil {
				charmlog.Error("Failed to read/write content to a file", "error", err)
			}
		}
	}
}
