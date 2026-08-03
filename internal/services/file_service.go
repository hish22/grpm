package services

/*
 * File service.
 *
 * Provide different actions to deal
 * with reading, and writing files and directors.
 */

import (
	"errors"
	"os"
	"strings"
)

type fileService struct {
	path *PathService
	mode os.FileMode
}

// Create new file service. FileService expect a PathService with a mode for the file.
func CreateFileService(path *PathService, mode os.FileMode) (*fileService, error) {
	if path.isEmpty() {
		return nil, errors.New("Failed to create FileService, PathService must not be empty")
	}
	return &fileService{
		path,
		mode,
	}, nil
}

// Change the FileService's PathService.
func (f *fileService) ChangePath(newPath *PathService) error {
	if newPath == nil {
		return errors.New("Failed to point to other path, NewPath argument is nil.")
	}
	f.path = newPath
	return nil
}

// Change the FileService's mode.
func (f *fileService) ChangeMode(newMode os.FileMode) {
	f.mode = newMode
}

// General file open.
func (f *fileService) Open(fileIndex int, flags int) (*os.File, error) {
	filePath, err := f.path.String(fileIndex)

	if err != nil {
		return nil, err
	}

	return os.OpenFile(filePath, flags, f.mode)
}

// Open file as read-only file.
func (f *fileService) OpenROFile(fileIndex int) (*os.File, error) {
	return f.Open(fileIndex, os.O_RDONLY)
}

// Create a file with RW flag. (TMP)
func (f *fileService) NewRWFile(fileIndex int) (*os.File, error) {
	return f.Open(fileIndex, os.O_CREATE|os.O_RDWR)
}

// Create a new file.
func (f *fileService) NewFile(fileIndex int) (*os.File, error) {
	return f.Open(fileIndex, os.O_CREATE)
}

// Create a set of new files. (could be refactored)
func (f *fileService) NewFiles() ([]*os.File, error) {
	filesLength := len(f.path.files)
	files := make([]*os.File, filesLength)
	for i := 0; i < len(f.path.files); i++ {
		file, err := f.NewFile(i)
		if err != nil {
			return nil, err
		}
		files = append(files, file)
	}
	return nil, nil
}

// Create a new temporary file.
func (f *fileService) NewTmpFile(fileIndex int) (*os.File, error) {
	file, err := f.path.GetFile(fileIndex)
	if err != nil {
		return nil, err
	}
	return os.CreateTemp(f.path.GetDirectoriesPath(), strings.Join([]string{file, "*"}, "-"))
}

// Create a set of new temporary files. (could be refactored)
func (f *fileService) NewTmpFiles() ([]*os.File, error) {
	filesLength := len(f.path.files)
	tmpFiles := make([]*os.File, filesLength)
	for i := 0; i < len(f.path.files); i++ {
		file, err := f.NewTmpFile(i)
		if err != nil {
			return nil, err
		}
		tmpFiles = append(tmpFiles, file)
	}
	return nil, nil
}

// Create a new set of directories. NewDir expect a inner boolean
// value, to make inner recursively directories or separated.
func (f *fileService) NewDir(inner bool) error {
	if inner == true {
		return os.MkdirAll(f.path.GetDirectoriesPath(), f.mode)
	} else {
		dirs := f.path.GetDirectories()
		for _, dir := range dirs {
			err := os.Mkdir(dir, f.mode)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// Read a stream of bytes from a file.
func (f *fileService) ReadStream(fileIndex int) ([]byte, error) {
	filePath, err := f.path.String(fileIndex)

	if err != nil {
		return nil, err
	}

	bytes, err := os.ReadFile(filePath)

	return bytes, err
}

// Read a stream of bytes from a file, and return it as a string.
func (f *fileService) ReadStreamAsString(fileIndex int) (string, error) {
	data, err := f.ReadStream(fileIndex)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// Write a stream of bytes to a file.
func (f *fileService) WriteStream(file *os.File, fileIndex int, stream []byte) error {
	filePath, err := f.path.String(fileIndex)
	if err != nil {
		return err
	}
	if file == nil {
		return os.WriteFile(filePath, stream, f.mode)
	} else {
		return os.WriteFile(file.Name(), stream, f.mode)
	}
}

// Write into intermediate temporary file before the regular file
func (f *fileService) ContingencySafeWrite(fileIndex int, stream []byte) error {
	tFile, err := f.NewTmpFile(fileIndex)

	if err != nil {
		return err
	}

	defer func() {
		os.Remove(tFile.Name())
	}()

	err = f.WriteStream(tFile, fileIndex, stream)
	if err != nil {
		tFile.Close()
		return err
	}

	err = tFile.Close()
	if err != nil {
		return err
	}

	filePath, err := f.path.String(fileIndex)

	if err != nil {
		return err
	}

	err = os.Rename(tFile.Name(), filePath)

	if err != nil {
		return err
	}

	return nil
}

// Provide the fileService mode.
func (f *fileService) GetMode() os.FileMode {
	return f.mode
}

// Provide a list of directories Entries.
func (f *fileService) DirsEntries() ([]os.DirEntry, error) {
	dirs, err := os.ReadDir(f.path.GetDirectoriesPath())

	if err != nil {
		return nil, err
	}

	return dirs, nil
}

// Provdie a list of directory Entries names.
func (f *fileService) DirsNames() ([]string, error) {
	dirs, err := f.DirsEntries()

	if err != nil {
		return nil, err
	}

	dirNames := make([]string, len(dirs))

	for _, dir := range dirs {
		dirNames = append(dirNames, dir.Name())
	}

	return dirNames, nil
}
