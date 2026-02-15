package setup

import (
	"archive/tar"
	"compress/gzip"
	corehttp "hish22/grpm/internal/coreHttp"
	"io"
	"os"

	charmlog "github.com/charmbracelet/log"
)

func unzipTar(header *tar.Header, file io.Reader) {
	switch header.Typeflag {

	case tar.TypeDir:
		err := os.MkdirAll(corehttp.WriteLibFilePath(header.Name), 0755)
		if err != nil {
			charmlog.Fatal("Failed to create directory", "error", err)
		}

	case tar.TypeReg:
		f, err := os.Create(corehttp.WriteLibFilePath(header.Name))
		if err != nil {
			charmlog.Fatal("Failed to create file", "error", err)
		}
		_, err = io.Copy(f, file)
		if err != nil {
			charmlog.Fatal("Failed to copy contents of a file", "error", err)
		}
	}
}

func tarReader(cfile io.Reader, location string) {
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
		charmlog.Info(header.Name)
		unzipTar(header, tarfile)
	}
}

func unzipFileTarGz(location string) {
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
	tarReader(gzip, location)
}
