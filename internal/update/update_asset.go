package update

import (
	"fmt"
	"hish22/grpm/internal/asset"
	"hish22/grpm/internal/install"
	"hish22/grpm/internal/release"
	"hish22/grpm/internal/structures"
	"regexp"
	"strconv"
	"strings"

	"github.com/charmbracelet/lipgloss"
	charmlog "github.com/charmbracelet/log"
)

var (
	assetNameStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#0000FF")).Bold(true)
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

func installUpdatedAsset(lr *structures.Release, oldAssetID int, version string) {
	ua := &structures.Assets{}
	for _, a := range lr.Assets {
		if a.AssetName == version {
			fmt.Println(a.AssetName)
			ua = &a
		}
	}
	asset.DeleteLastTrackedAssetById(oldAssetID)
	install.InstallSelectedAsset(version, ua, lr, false)
}

func UpdateToLatestAsset(repo string) {
	a := asset.FetchSpecificAsset(repo)
	latestA := release.FetchLatestRelease(repo)
	rx, err := regexp.Compile(`\d+.\d+.\d+`)
	if err != nil {
		charmlog.Fatal("Regex Failed to compile, ", "error", err)
	}

	b := rx.Find([]byte(a.Tag))
	lb := rx.Find([]byte(latestA.TagName))

	// Asset Tag
	version := strings.Split(string(b), ".")
	major, err := strconv.Atoi(version[0])
	minor, err := strconv.Atoi(version[1])
	patch, err := strconv.Atoi(version[2])

	if err != nil {
		charmlog.Fatal("Can't convert version to number", "error", err)
	}

	// Latest release Tag
	latestVersion := strings.Split(string(lb), ".")
	lmajor, err := strconv.Atoi(latestVersion[0])
	lminor, err := strconv.Atoi(latestVersion[1])
	lpatch, err := strconv.Atoi(latestVersion[2])

	if err != nil {
		charmlog.Fatal("Can't convert latest version to number", "error", err)
	}
	// Replace if new version found/or nothing changes
	newVersion := updateVersion(a.AssetName, string(lb))
	isUpdateable := false
	if lmajor > major {
		charmlog.Info("Major Updating...")
		isUpdateable = true
	} else if lminor > minor {
		charmlog.Info("Minor Updating...")
		isUpdateable = true
	} else if lpatch > patch {
		charmlog.Info("Patch Updating...")
		isUpdateable = true
	} else {
		charmlog.Info(assetNameStyle.Render(a.AssetName) + " installed with its latest version.")
	}

	if isUpdateable {
		installUpdatedAsset(latestA, a.ID, newVersion)
	}

}
