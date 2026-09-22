package services

/*
 * File service.
 *
 * Provide different actions to deal
 * with reading, and writing files and directors.
 */

// NEEDS REFACTOR
// SOME FUNCTIONS (MUST,MAY) BE REMOVED.
// SOME FUNCTIONS (MUST,MAY) BE COMPOUND WITH OTHER FUNCTION(S).
// SOME FUNCTIONS (MUST,MAY) CHECKED.

import (
	"errors"
	"io"
	"os"
	"strings"
)

type fileService struct {
	path  *PathService
	index int
	mode  os.FileMode
}

// Create new file service. FileService expect a PathService with a mode for the file.
func CreateFileService(path *PathService, mode os.FileMode) (*fileService, error) {
	if path.isEmpty() {
		return nil, errors.New("Failed to create FileService, PathService must not be empty")
	}
	return &fileService{
		path,
		0, // Default with index 0
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

// Change the FileService's Current Index.
func (f *fileService) ChangeIndex(fileIndex int) error {
	if len(f.path.files) < fileIndex {
		return errors.New("Failed to change file index, wrong index provided.")
	}

	f.index = fileIndex
	return nil
}

// Change the FileService's mode.
func (f *fileService) ChangeMode(newMode os.FileMode) {
	f.mode = newMode
}

// General file open.
func (f *fileService) Open(flags int) (*os.File, error) {
	filePath, err := f.path.String(f.index)

	if err != nil {
		return nil, err
	}

	return os.OpenFile(filePath, flags, f.mode)
}

// Open file as read-only. (TMP)
func (f *fileService) OpenROFile() (*os.File, error) {
	return f.Open(os.O_RDONLY)
}

// Open file as write-only. (TMP)
func (f *fileService) OpenWOFile() (*os.File, error) {
	return f.Open(os.O_WRONLY)
}

// Create a file with RW flag. (TMP)
func (f *fileService) NewRWFile() (*os.File, error) {
	return f.Open(os.O_CREATE | os.O_RDWR)
}

// Create a new file. (TMP)
func (f *fileService) NewFile() (*os.File, error) {
	return f.Open(os.O_CREATE)
}

// Create a set of new files. (could be refactored)
func (f *fileService) NewFiles() ([]*os.File, error) {
	filesLength := len(f.path.files)
	files := make([]*os.File, filesLength)
	for i := 0; i < len(f.path.files); i++ {
		f.ChangeIndex(i)
		file, err := f.NewFile()
		if err != nil {
			return nil, err
		}
		files = append(files, file)
	}
	return nil, nil
}

// Create a new temporary file.
func (f *fileService) NewTmpFile() (*os.File, error) {
	file, err := f.path.GetFile(f.index)
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
		f.ChangeIndex(i)
		file, err := f.NewTmpFile()
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
func (f *fileService) ReadStream() ([]byte, error) {
	filePath, err := f.path.String(f.index)

	if err != nil {
		return nil, err
	}

	bytes, err := os.ReadFile(filePath)

	return bytes, err
}

// Read a stream of bytes from a file, and return it as a string.
func (f *fileService) ReadStreamAsString() (string, error) {
	data, err := f.ReadStream()
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// Write a stream of bytes to a file.
func (f *fileService) WriteStream(file *os.File, stream []byte) error {
	filePath, err := f.path.String(f.index)
	if err != nil {
		return err
	}
	if file == nil {
		return os.WriteFile(filePath, stream, f.mode)
	} else {
		return os.WriteFile(file.Name(), stream, f.mode)
	}
}

// Provide ability to copy content from a read stream to a write stream.
func (f *fileService) CopyFile(stream io.Reader) error {
	file, err := f.OpenWOFile()

	if err != nil {
		return err
	}

	_, err = io.Copy(file, stream)

	return err
}

// Copy into intermediate temporary file before the regular file
func (f *fileService) ContingencySafeCopy(stream io.Reader) error {
	tFile, err := f.NewTmpFile()

	if err != nil {
		return err
	}

	defer func() {
		tFile.Close()
		os.Remove(tFile.Name())
	}()

	_, err = io.Copy(tFile, stream)

	if err != nil {
		tFile.Close()
		return err
	}

	filePath, err := f.path.String(f.index)

	if err != nil {
		return err
	}

	err = os.Rename(tFile.Name(), filePath)

	if err != nil {
		return err
	}

	return nil

}

// Write into intermediate temporary file before the regular file
func (f *fileService) ContingencySafeWrite(stream []byte) error {
	tFile, err := f.NewTmpFile()

	if err != nil {
		return err
	}

	defer func() {
		tFile.Close()
		os.Remove(tFile.Name())
	}()

	err = f.WriteStream(tFile, stream)
	if err != nil {
		tFile.Close()
		return err
	}

	err = tFile.Close()
	if err != nil {
		return err
	}

	filePath, err := f.path.String(f.index)

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

func (f *fileService) GetIndex() int {
	return f.index
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
