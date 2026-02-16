package link

import (
	"hish22/grpm/internal/config"

	charmlog "github.com/charmbracelet/log"
)

func WriteDownloadsFilePath(filename string) string {
	downlaodsPath, err := config.GrpmDownloadedDirPath()
	if err != nil {
		charmlog.Error("Failed to return download path", "error", err)
	}
	downlaodsPath.Asset = filename
	return downlaodsPath.String()
}

func WriteLibFilePath(filename string) string {
	libPath, err := config.GrpmLibraryDirPath()
	if err != nil {
		charmlog.Error("Failed to return library path", "error", err)
	}
	libPath.Asset = filename
	return libPath.String()
}
