package setup

import (
	"archive/tar"
	"hish22/grpm/internal/link"
	"io"
	"os"
	"path/filepath"

	charmlog "github.com/charmbracelet/log"
)

func unzipTar(_ string, header *tar.Header, file io.Reader, _ int) {
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
	}
}

func tarReader(_ string, cfile io.Reader, location string, assetID int) {
	// read archive .tar file
	tarfile := tar.NewReader(cfile)
	charmlog.Info("Extracting..", "asset", location)
	for {
		header, err := tarfile.Next()
		if err == io.EOF {
			charmlog.Warn("EOF reached")
			break
		}
		if err != nil {
			charmlog.Error("Failed to read tar", "error", err)
		}
		charmlog.Info(header.Name, "asset_id", assetID)
		unzipTar("", header, tarfile, assetID)
	}
}
