package util

import "regexp"

// UpdateVersion replaces a SemVer string within a filename automatically
func UpdateVersion(filename, newVersion string) string {
	// Pattern breakdown:
	// ^(.*?-)          : Group 1 - Matches everything from the start up to the last dash before version
	// [0-9]+\.[0-9]+\.[0-9]+ : Matches the actual version (e.g., 1.1.3)
	// (-.*|.*)$        : Group 2 - Matches the rest of the string (extension, arch, etc.)
	re := regexp.MustCompile(`^(.*?)[0-9]+\.[0-9]+\.[0-9]+(.*)$`)

	// ${1} is the prefix, ${2} is the suffix
	return re.ReplaceAllString(filename, "${1}"+newVersion+"${2}")
}
