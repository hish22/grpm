package setup

import (
	"archive/tar"
	"hish22/grpm/internal/asset"
	"hish22/grpm/internal/link"
	"io"
	"os"
	"path/filepath"
	"strings"

	charmlog "github.com/charmbracelet/log"
)

func unzipTar(repo string, header *tar.Header, assetName string, file io.Reader, assetID int, force bool) {
	switch header.Typeflag {

	case tar.TypeDir:
		err := os.MkdirAll(assetName, 0755)
		if err != nil {
			charmlog.Fatal("Failed to create directory", "error", err)
		}

	case tar.TypeReg:
		// asssetlink := link.WriteLibFilePath(header.Name)
		// Ensure parent directory exists
		// if err := os.MkdirAll(link.WriteLibFilePath(filepath.Dir(header.Name)), 0755); err != nil {
		// 	charmlog.Warn("Failed to create directory", "error", err)
		// }
		f, err := os.Create(assetName)
		if err != nil {
			charmlog.Fatal("Failed to create file", "error", err)
		}
		_, err = io.Copy(f, file)
		if err != nil {
			charmlog.Fatal("Failed to copy contents of a file", "error", err)
		}
		binaryName := strings.Split(assetName, "/")
		SymlinkAsset(repo, assetName, binaryName[len(binaryName)-1], assetID, force)
	}
}

func tarReader(repo string, cfile io.Reader, location string, assetID int, force bool) {
	// read archive .tar file
	tarfile := tar.NewReader(cfile)
	charmlog.Info("Extracting..", "asset", location)
	// isSetupFileRegisterd := false
	isThereAParentDir := false
	parentPath := ""
	for {
		header, err := tarfile.Next()
		if err == io.EOF {
			charmlog.Warn("EOF reached")
			break
		}
		if err != nil {
			charmlog.Error("Failed to read tar", "error", err)
		}
		// check if first file is parent dir
		if header.FileInfo().IsDir() && !isThereAParentDir {
			isThereAParentDir = true
			asset.InsertFileSetupLocation(link.WriteLibFilePath(header.Name), assetID)
			// if not create one
		} else if !isThereAParentDir {
			parentPath = link.WriteLibFilePath(repo)
			if err := os.MkdirAll(parentPath, 0755); err != nil {
				charmlog.Warn("Failed to create directory", "error", err)
			}
			isThereAParentDir = true
			asset.InsertFileSetupLocation(parentPath, assetID)
		}
		assetName := filepath.Join(parentPath, header.Name)
		// if !isSetupFileRegisterd && isThereAParentDir {

		// 	isSetupFileRegisterd = true
		// }

		charmlog.Info(assetName, "asset_id", assetID)
		unzipTar(repo, header, assetName, tarfile, assetID, force)
	}
}
