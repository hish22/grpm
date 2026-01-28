package util

import "regexp"

func CleanHtmlTags(value string) string {
	rgx := regexp.MustCompile("<[^>]*>")
	return rgx.ReplaceAllString(value, "")
}
