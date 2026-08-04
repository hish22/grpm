package services

import (
	"fmt"
	"strings"
	"testing"
)

func TestFileServiceChangePath(t *testing.T) {
	path := CreatePathService([]string{"help"}, nil)

	file, err := CreateFileService(path, 0755)

	if err != nil {
		t.Error(err)
	}

	fmt.Println(file.DirsNames())

	newPath := CreatePathService([]string{"testfile"}, []string{"welcome.txt"})

	err = file.ChangePath(newPath)

	if err != nil {
		t.Error(err)
	}

	fmt.Println(file.DirsNames())
}

func TestFileServiceChangeMode(t *testing.T) {
	path := CreatePathService([]string{"testfile"}, []string{"no.txt"})

	file, err := CreateFileService(path, 0755)

	if err != nil {
		t.Error(err)
	}

	fmt.Println(file.GetMode().String())

	file.ChangeMode(0644)

	fmt.Println(file.GetMode().String())

}

func TestFileServiceNewFiles(t *testing.T) {
	path := CreatePathService([]string{"testfile"}, []string{"no.txt", "yes.txt", "aha.txt"})

	file, err := CreateFileService(path, 0755)

	if err != nil {
		t.Error(err)
	}

	files, err := file.NewFiles()

	defer func() {
		for _, file := range files {
			file.Close()
		}
	}()

	if err != nil {
		t.Error(err)
	}

}

func TestFileServiceNewTmpFiles(t *testing.T) {
	path := CreatePathService([]string{"testfile"}, []string{"no.txt", "yes.txt", "aha.txt"})

	file, err := CreateFileService(path, 0755)

	if err != nil {
		t.Error(err)
	}

	files, err := file.NewTmpFiles()

	defer func() {
		for _, file := range files {
			file.Close()
		}
	}()

	if err != nil {
		t.Error(err)
	}

}

func TestFileServiceNewDirs(t *testing.T) {
	path := CreatePathService([]string{"testfile", "files", "boo"}, []string{})

	file, err := CreateFileService(path, 0755)

	if err != nil {
		t.Error(err)
	}

	err = file.NewDir(true)

	if err != nil {
		t.Error(err)
	}

}

func TestFileServiceReadStream(t *testing.T) {
	path := CreatePathService([]string{""}, []string{"content.txt"})

	file, err := CreateFileService(path, 0755)

	if err != nil {
		t.Error(err)
	}

	data, err := file.ReadStreamAsString()

	if err != nil {
		t.Error(err)
	}

	fmt.Println(data)

}

func TestFileServiceContingencySafeWrite(t *testing.T) {
	path := CreatePathService([]string{""}, []string{"test.txt", "cache.txt"})

	file, err := CreateFileService(path, 0644)

	if err != nil {
		t.Error(err)
		return
	}

	file.ContingencySafeWrite([]byte("Welcome to my world!"))
	file.ChangeIndex(1)
	file.ContingencySafeWrite([]byte("Welcome to my cahce file!"))
}

func TestFileServiceCopyFile(t *testing.T) {
	path := CreatePathService([]string{""}, []string{"content.txt"})

	file, err := CreateFileService(path, 0755)

	if err != nil {
		t.Error(err)
	}

	stringReader := strings.NewReader("Welcome to copy file test")

	err = file.CopyFile(stringReader)

	if err != nil {
		t.Error(err)
	}
}

func TestFileServiceContingencySafeCopy(t *testing.T) {
	path := CreatePathService([]string{""}, []string{"content.txt"})

	file, err := CreateFileService(path, 0755)

	if err != nil {
		t.Error(err)
	}

	stringReader := strings.NewReader("Welcome to copy file test")
	err = file.ContingencySafeCopy(stringReader)

	if err != nil {
		t.Error(err)
	}
}
