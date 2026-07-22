package services

import (
	"fmt"
	"testing"
)

func TestPathService(t *testing.T) {
	path, err := CreatePathService([]string{"root", "home", "Desktop", ".cache"}, "temp.json")

	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println(path.String())
	fmt.Println(path.GetDirectories())
	fmt.Println(path.FileExt(false))

}

func TestPathServiceWithEmptyDirectories(t *testing.T) {
	path, err := CreatePathService([]string{}, "temp.json")

	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println(path.String())

}
