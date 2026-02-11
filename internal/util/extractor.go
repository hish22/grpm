package util

import (
	"regexp"

	charmlog "github.com/charmbracelet/log"
)

func removeDot(extensions []string) {
	for i := range len(extensions) {
		extensions[i] = extensions[i][1:]
	}
}

func ExtensionExtractor(file string) []string {
	regx, err := regexp.Compile(`\.\D\w*`)
	if err != nil {
		charmlog.Fatal("Failed to compile extractor regular expression", "error", err)
	}
	extensions := regx.FindAllString(file, -1)
	removeDot(extensions)
	return extensions
}
