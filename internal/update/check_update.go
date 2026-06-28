package update

import (
	"hish22/grpm/internal/structures"

	charmlog "github.com/charmbracelet/log"
)

func CheckUpdate(repo string) (*structures.TrackedAsset, *structures.Release, []byte, []byte, bool) {
	// Fetch Current and latest assets
	currentAsset, latestAsset := fetchReleases(repo)

	// Build regex
	rx := buildregx()
	b := rx.Find([]byte(currentAsset.Tag))
	lb := rx.Find([]byte(latestAsset.TagName))

	// Asset Tag
	major, minor, patch := extractVersionSet(b)
	// Latest release Tag
	lmajor, lminor, lpatch := extractVersionSet(lb)

	checkStatus := isUpdateable(currentAsset.AssetName, lmajor, major, lminor, minor, lpatch, patch)

	if checkStatus {
		charmlog.Info("This asset has a new version.", "Current Version", currentAsset.Tag, "New Version", latestAsset.TagName)
	}
	return currentAsset, latestAsset, b, lb, checkStatus
}
