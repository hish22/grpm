package services

import (
	"fmt"
	"testing"
)

func TestPathService(t *testing.T) {
	path := CreatePathService([]string{"root", "home", "Desktop", ".cache"}, []string{"temp.json", "about.json"})

	pathForFile0, err := path.String(0)

	if err != nil {
		t.Error(err)
	}

	pathForFile1, err := path.String(1)

	if err != nil {
		t.Error(err)
	}

	fmt.Println(pathForFile0)
	fmt.Println(pathForFile1)

	fmt.Println(path.GetDirectories())
	fmt.Println(path.GetDirectoriesPath())

	ext, err := path.FileExt(0, false)

	if err != nil {
		t.Error(err)
	}

	fmt.Println(ext)

}

func TestPathServiceWithEmptyDirectories(t *testing.T) {
	path := CreatePathService([]string{}, []string{"temp.json"})

	fmt.Println(path.String(0))

}
