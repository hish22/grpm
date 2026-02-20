package update

import (
	"hish22/grpm/internal/asset"
	"hish22/grpm/internal/install"
	"hish22/grpm/internal/release"
	"hish22/grpm/internal/remove"
	"hish22/grpm/internal/structures"
	"hish22/grpm/internal/util"
	"regexp"
	"strconv"
	"strings"

	"github.com/charmbracelet/lipgloss"
	charmlog "github.com/charmbracelet/log"
)

var (
	assetNameStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#0000FF")).Bold(true)
)

func installUpdatedAsset(repo string, lr *structures.Release, oldAsset *structures.TrackedAsset, version string) {
	// Check if user is running this with privileged execution
	if !util.IsAdministrator() {
		charmlog.Error("Please run this command with privilege execution mode")
		return
	}

	ua := &structures.Assets{}
	for _, a := range lr.Assets {
		if a.AssetName == version {
			ua = &a
		}
	}
	setupStatus := asset.AssetSetupTrackStatus(oldAsset.ID)
	remove.RemoveAssetByID(oldAsset.ID, oldAsset.Location)
	install.InstallSelectedAsset(repo, ua, lr, setupStatus)
}

func buildregx() *regexp.Regexp {
	rx, err := regexp.Compile(`\d+.\d+.\d+`)
	if err != nil {
		charmlog.Fatal("Regex Failed to compile, ", "error", err)
	}
	return rx
}

func extractVersionSet(tag []byte) (int, int, int) {
	version := strings.Split(string(tag), ".")
	major, err := strconv.Atoi(version[0])
	minor, err := strconv.Atoi(version[1])
	patch, err := strconv.Atoi(version[2])
	if err != nil {
		charmlog.Fatal("Can't convert latest version to number", "error", err)
	}
	return major, minor, patch
}

func UpdateToLatestAsset(repo string) {
	// Fetch Specific asset
	a, err := asset.FetchSpecificAsset(repo)
	if err != nil {
		charmlog.Fatal("Failed to fetch specified repository", "repo", repo, "error", err)
	}
	// Fetch latest repo release
	latestA, err := release.FetchLatestRelease(repo)
	if err != nil {
		charmlog.Fatal("Failed to fetch latest release", "repo", repo, "error", err)
	}
	// Build regex
	rx := buildregx()
	b := rx.Find([]byte(a.Tag))
	lb := rx.Find([]byte(latestA.TagName))

	// Asset Tag
	major, minor, patch := extractVersionSet(b)
	// Latest release Tag
	lmajor, lminor, lpatch := extractVersionSet(lb)

	// Replace if new version found/or nothing changes
	newVersion := util.UpdateVersion(a.AssetName, string(lb))
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
		installUpdatedAsset(repo, latestA, &a, newVersion)
	}

}
