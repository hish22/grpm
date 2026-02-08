package update

import (
	"fmt"
	"hish22/grpm/internal/asset"
	"hish22/grpm/internal/install"
	"hish22/grpm/internal/persistance"
	"hish22/grpm/internal/release"
	"hish22/grpm/internal/structures"
	"log"
	"regexp"
	"strconv"
	"strings"

	charmlog "github.com/charmbracelet/log"
)

// UpdateVersion replaces a SemVer string within a filename automatically
func updateVersion(filename, newVersion string) string {
	// Pattern breakdown:
	// ^(.*?-)          : Group 1 - Matches everything from the start up to the last dash before version
	// [0-9]+\.[0-9]+\.[0-9]+ : Matches the actual version (e.g., 1.1.3)
	// (-.*|.*)$        : Group 2 - Matches the rest of the string (extension, arch, etc.)
	re := regexp.MustCompile(`^(.*?)[0-9]+\.[0-9]+\.[0-9]+(.*)$`)

	// ${1} is the prefix, ${2} is the suffix
	return re.ReplaceAllString(filename, "${1}"+newVersion+"${2}")
}

func deleteLastTrackedAsset(id int) {
	db := persistance.OpenMetadataDB()
	_, err := db.Exec("DELETE FROM asset WHERE id=?", id)
	if err != nil {
		charmlog.Fatal("Failed to delete last tracked asset", "error", err)
	}
}

func installUpdatedAsset(lr *structures.Release, oldAssetID int, version *string) {
	ua := &structures.Assets{}
	for _, a := range lr.Assets {
		if a.AssetName == *version {
			fmt.Println(a.AssetName)
			ua = &a
		}
	}
	deleteLastTrackedAsset(oldAssetID)
	install.InstallSelectedAsset(version, ua, lr)
}

func UpdateToLatestAsset(repo *string) {
	a := asset.FetchSpecificAsset(repo)
	latestA := release.FetchLatestRelease(repo)
	rx, err := regexp.Compile(`\d.\d.\d`)
	if err != nil {
		log.Fatal("Regex Failed to compile, ", err)
	}

	b := rx.Find([]byte(a.Tag))
	lb := rx.Find([]byte(latestA.TagName))
	// Asset Tag
	version := strings.Split(string(b), ".")
	major, err := strconv.Atoi(version[0])
	minor, err := strconv.Atoi(version[1])
	patch, err := strconv.Atoi(version[2])

	if err != nil {
		log.Fatal("Can't convert version to number", err)
	}

	// Latest release Tag
	latestVersion := strings.Split(string(lb), ".")
	lmajor, err := strconv.Atoi(latestVersion[0])
	lminor, err := strconv.Atoi(latestVersion[1])
	lpatch, err := strconv.Atoi(latestVersion[2])

	if err != nil {
		log.Fatal("Can't convert latest version to number", err)
	}
	// Replace if new version found/or nothing changes
	newVersion := updateVersion(a.AssetName, string(lb))

	if lmajor > major {
		charmlog.Info("Major Updating...")
		installUpdatedAsset(latestA, a.ID, &newVersion)
	} else if lminor > minor {
		charmlog.Info("Minor Updating...")
		installUpdatedAsset(latestA, a.ID, &newVersion)
	} else if lpatch > patch {
		charmlog.Info("Patch Updating...")
		installUpdatedAsset(latestA, a.ID, &newVersion)
	}

}
