package services

import (
	"fmt"
	"testing"
	"time"
)

func TestDownloadServiceBardown(t *testing.T) {
	link, err := CreateRequestLink(
		"https://github.com/junegunn/fzf/releases/download/v0.74.2/fzf-0.74.2-linux_amd64.tar.gz",
		[]string{}, []string{},
	)

	if err != nil {
		t.Error(err)
	}

	api, err := CreateApiRequest(link,
		map[string]string{},
		"GET",
		24*time.Hour)

	if err != nil {
		t.Error(err)
	}

	path := CreatePathService([]string{}, []string{"fzf-0.74.2-linux_amd64.tar.gz"})

	file, err := CreateFileService(path, 0755)

	if err != nil {
		t.Error(err)
	}

	download, err := CreateDownloadService(api, file)

	if err != nil {
		t.Error(err)
	}

	err = download.bardown()

	if err != nil {
		t.Error(err)
	}

	fmt.Println("File downloaded!")
}

func TestDownloadServiceDown(t *testing.T) {

	link, err := CreateRequestLink(
		"https://github.com/junegunn/fzf/releases/download/v0.74.2/fzf-0.74.2-linux_amd64.tar.gz",
		[]string{}, []string{},
	)

	if err != nil {
		t.Error(err)
	}

	api, err := CreateApiRequest(link,
		map[string]string{},
		"GET",
		24*time.Hour)

	if err != nil {
		t.Error(err)
	}

	path := CreatePathService([]string{}, []string{"fzf-0.74.2-linux_amd64.tar.gz"})

	file, err := CreateFileService(path, 0755)

	if err != nil {
		t.Error(err)
	}

	download, err := CreateDownloadService(api, file)

	if err != nil {
		t.Error(err)
	}

	err = download.down()

	if err != nil {
		t.Error(err)
	}

	fmt.Println("File downloaded!")

}
