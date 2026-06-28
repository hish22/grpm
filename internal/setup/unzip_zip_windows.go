package setup

import (
	"archive/zip"
	"hish22/grpm/internal/asset"
	"hish22/grpm/internal/config"
	"hish22/grpm/internal/util"
	"io"
	"os"
	"path/filepath"
	"strings"

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

	parentDir := libLink.String()
	parentDirFromZip := ""

	// First file was Dir, then create as parent dir
	isNotAparentDir := false

	// status of env vars
	isAssigned := false

	// Find parent directory
	for _, f := range zipf.File {
		// if the unziped file doesn't have parent dir
		// then we need to create one
		if !f.FileInfo().IsDir() && !isNotAparentDir {
			// open the .zip file
			zFileInfo, err := os.Open(location)
			if err != nil {
				charmlog.Error("Failed to open file", "error", err, "file", location)
			}
			// Extract the file name without the ext
			name := util.NameAndExtensionExtractor(filepath.Base(zFileInfo.Name()))[1]

			// create file name
			parentDir = filepath.Join(libLink.String(), name)

			// track setup location
			asset.InsertFileSetupLocation(parentDir, assetID)

			// Create the parent directory
			err = os.MkdirAll(parentDir, 0755)
			if err != nil {
				charmlog.Error("Failed to create direcotry", "error", err)
			}
			isNotAparentDir = true
			break
		} else {
			// track setup location
			parentDirFromZip = filepath.Join(parentDir, f.Name)
			asset.InsertFileSetupLocation(parentDirFromZip, assetID)
			break
		}
	}

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
		// Add environment variables if there is a /bin dir
		if strings.Contains(f.Name, "bin") && !isAssigned {
			RegisterEnvVar(assetPath)
			asset.InsertSymlinkOrEnvLocation(assetPath, assetID)
			isAssigned = true
		}
	}
	// if there is no a /bin dir, register the parent dir
	if !isAssigned && isNotAparentDir {
		RegisterEnvVar(parentDir)
		asset.InsertSymlinkOrEnvLocation(parentDir, assetID)
	} else if !isAssigned && !isNotAparentDir {
		RegisterEnvVar(parentDirFromZip)
		asset.InsertSymlinkOrEnvLocation(parentDirFromZip, assetID)
	}

}
