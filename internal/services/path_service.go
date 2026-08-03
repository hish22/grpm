package services

/*
 * Path service.
 *
 * Provide ability to create
 * and manage path strings.
 */

import (
	"errors"
	"path/filepath"
	"regexp"
	"strings"
)

type PathService struct {
	directories []string
	files       []string
}

// Create new Path service. PathService expect a list of directories and a file name.
func CreatePathService(directories []string, files []string) *PathService {
	return &PathService{
		directories,
		files,
	}

}

// Change the value of (p.files)
func (p PathService) ChangeFileName(files []string) {
	p.files = files
}

// Change the value of (p.directories).
func (p PathService) ChangeDirectories(directories []string) {
	p.directories = directories
}

// Check if PathService object is empty.
func (p PathService) isEmpty() bool {
	return len(p.directories) == 0 && len(p.files) == 0
}

// Provide a string version of the path, String expect file index.
func (p PathService) String(i int) (string, error) {
	if len(p.files) <= i {
		return "", errors.New("Failed to provide a string, wrong index")
	}

	dirsPath := filepath.Join(p.directories...)

	return filepath.Join(dirsPath, p.files[i]), nil
}

// Provide a file access as []string.
func (p PathService) GetFiles() []string {
	return p.files
}

// Provide a specific file string.
func (p PathService) GetFile(i int) (string, error) {
	if len(p.files) <= i {
		return "", errors.New("Failed to get specified file, wrong index")
	}
	return p.files[i], nil
}

// Provide directories access as []string.
func (p PathService) GetDirectories() []string {
	return p.directories
}

// Provide a string version of directories.
func (p PathService) GetDirectoriesPath() string {
	return filepath.Join(p.GetDirectories()...)
}

// Provide a file extention.
func (p PathService) FileExt(i int, nodot bool) (string, error) {
	if len(p.files) <= i {
		return "", errors.New("Failed to extract file extention, wrong index")
	}

	if nodot {
		_, a, found := strings.Cut(filepath.Ext(p.files[i]), ".")
		if found {
			return a, nil
		}
		return "", nil
	}
	return filepath.Ext(p.files[i]), nil
}

// Provide a file without extention.
func (p PathService) RmvExt(i int) (string, error) {
	if len(p.files) <= i {
		return "", errors.New("Failed to remove file extention, wrong index")
	}

	re := regexp.MustCompile(`^(\w)+`)
	m := re.FindString(p.files[i])
	return m, nil
}
