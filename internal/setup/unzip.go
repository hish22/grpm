package setup

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"os"

	charmlog "github.com/charmbracelet/log"
)

// Create a switch for different file extensions
// based on the extensions, we perfom the setup phase

func UnzipFileTarGz(location string) {
	// open the compressed file
	file, err := os.Open(location)
	if err != nil {
		charmlog.Fatal("Failed to open compressed file", "error", err)
	}
	defer file.Close()
	// Uncompress the file of .gz
	gzip, err := gzip.NewReader(file)
	if err != nil {
		charmlog.Fatal("Failed to uncompress file", "error", err)
	}
	defer gzip.Close()
	// read archive .tar file
	tarfile := tar.NewReader(gzip)
	charmlog.Info("Extracting..", "asset", location)
	for {
		header, err := tarfile.Next()
		if err != nil {
			charmlog.Warn("Failed to read tar", "error", err)
		}
		if err == io.EOF {
			break
		}
		charmlog.Info(header.Name)

		switch header.Typeflag {

		case tar.TypeDir:
			err = os.MkdirAll(header.Name, 0755)
			if err != nil {
				charmlog.Fatal("Failed to create directory", "error", err)
			}

		case tar.TypeReg:
			f, err := os.Create(header.Name)
			if err != nil {
				charmlog.Fatal("Failed to create file", "error", err)
			}
			_, err = io.Copy(f, tarfile)
			if err != nil {
				charmlog.Fatal("Failed to copy contents of a file", "error", err)
			}
		}
	}
}
