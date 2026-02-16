package setup

import (
	"archive/tar"
	"compress/gzip"
	"hish22/grpm/internal/asset"
	"hish22/grpm/internal/link"
	"io"
	"os"
	"strings"

	charmlog "github.com/charmbracelet/log"
)

func unzipTar(header *tar.Header, file io.Reader, assetID int) {
	switch header.Typeflag {

	case tar.TypeDir:
		err := os.MkdirAll(link.WriteLibFilePath(header.Name), 0755)
		if err != nil {
			charmlog.Fatal("Failed to create directory", "error", err)
		}

	case tar.TypeReg:
		link := link.WriteLibFilePath(header.Name)
		f, err := os.Create(link)
		if err != nil {
			charmlog.Fatal("Failed to create file", "error", err)
		}
		_, err = io.Copy(f, file)
		if err != nil {
			charmlog.Fatal("Failed to copy contents of a file", "error", err)
		}
		binaryName := strings.Split(header.Name, "/")
		SymlinkAsset(link, binaryName[len(binaryName)-1], assetID)
	}
}

func tarReader(cfile io.Reader, location string, assetID int) {
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
		charmlog.Info(header.Name)
		unzipTar(header, tarfile, assetID)
	}
}

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
