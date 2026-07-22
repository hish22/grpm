package services

import (
	"errors"
	"path/filepath"
	"strings"
)

/*
 * Path service.
 *
 * Provide ability to create
 * and manage path strings.
 */

type PathService struct {
	directories []string
	file        string
}

// Create new Path service. PathService expect a list of directories and a file name.
func CreatePathService(directories []string, file string) (*PathService, error) {
	if len(directories) == 0 {
		return nil, errors.New("Failed to create PathService, directories must not be of length 0")
	}

	return &PathService{
		directories,
		file,
	}, nil

}

// Provide a string version of the path.
func (p PathService) String() string {
	dirsPath := filepath.Join(p.directories...)
	return filepath.Join(dirsPath, p.file)
}

// Provide a file access as string.
func (p PathService) GetFile() string {
	return p.file
}

// Provide directories access as string path.
func (p PathService) GetDirectories() string {
	return filepath.Join(p.directories...)
}

// Provide a file extention.
func (p PathService) FileExt(nodot bool) string {
	if nodot {
		_, a, found := strings.Cut(filepath.Ext(p.file), ".")
		if found {
			return a
		}
		return ""
	}
	return filepath.Ext(p.file)
}
