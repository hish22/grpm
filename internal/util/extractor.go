package util

import (
	"regexp"
	"strings"

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
		charmlog.Error("Failed to compile extractor regular expression", "error", err)
	}
	extensions := regx.FindAllString(file, -1)
	removeDot(extensions)
	return extensions
}

func NameAndExtensionExtractor(file string) []string {
	regx, err := regexp.Compile(`(.+)\.\D\w*`)
	if err != nil {
		charmlog.Error("Failed to compile name and ext extractor regular expression", "error", err)
	}
	namesAndExtensions := regx.FindStringSubmatch(file)
	return namesAndExtensions
}

// owner/repo -> repo
func RepoNameExtractor(repo string) string {
	if strings.Contains(repo, "/") {
		return strings.Split(repo, "/")[1]
	}
	return repo
}
