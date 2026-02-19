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

func unzipTar(repo string, header *tar.Header, file io.Reader, assetID int) {
	switch header.Typeflag {

	case tar.TypeDir:
		err := os.MkdirAll(link.WriteLibFilePath(header.Name), 0755)
		if err != nil {
			charmlog.Fatal("Failed to create directory", "error", err)
		}

	case tar.TypeReg:
		asssetlink := link.WriteLibFilePath(header.Name)
		// Ensure parent directory exists
		if err := os.MkdirAll(link.WriteLibFilePath(filepath.Dir(header.Name)), 0755); err != nil {
			charmlog.Warn("Failed to create directory", "error", err)
		}
		f, err := os.Create(asssetlink)
		if err != nil {
			charmlog.Fatal("Failed to create file", "error", err)
		}
		_, err = io.Copy(f, file)
		if err != nil {
			charmlog.Fatal("Failed to copy contents of a file", "error", err)
		}
		binaryName := strings.Split(header.Name, "/")
		SymlinkAsset(repo, asssetlink, binaryName[len(binaryName)-1], assetID)
	}
}

func tarReader(repo string, cfile io.Reader, location string, assetID int) {
	// read archive .tar file
	tarfile := tar.NewReader(cfile)
	charmlog.Info("Extracting..", "asset", location)
	isSetupFileRegisterd := false
	for {
		header, err := tarfile.Next()
		if err == io.EOF {
			charmlog.Warn("EOF reached")
			break
		}
		if err != nil {
			charmlog.Error("Failed to read tar", "error", err)
		}
		if !isSetupFileRegisterd {
			asset.InsertFileSetupLocation(link.WriteLibFilePath(header.Name), assetID)
			isSetupFileRegisterd = true
		}
		charmlog.Info(header.Name, "asset_id", assetID)
		unzipTar(repo, header, tarfile, assetID)
	}
}
