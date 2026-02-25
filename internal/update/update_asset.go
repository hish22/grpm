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

func installUpdatedAsset(repo string, lr *structures.Release, oldAsset *structures.TrackedAsset, version string, force bool) {
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
	install.InstallSelectedAsset(repo, ua, lr, setupStatus, force)
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

func isUpdateable(assetName string, lmajor int, major int, lminor int, minor int, lpatch int, patch int) bool {
	if lmajor > major {
		charmlog.Info("Major Updating...")
		return true
	} else if lminor > minor {
		charmlog.Info("Minor Updating...")
		return true
	} else if lpatch > patch {
		charmlog.Info("Patch Updating...")
		return true
	} else {
		charmlog.Info(assetNameStyle.Render(assetName) + " installed with its latest version.")
		return false
	}
}

func fetchReleases(repo string) (*structures.TrackedAsset, *structures.Release) {
	// Fetch Specific asset
	a, err := asset.FetchSpecificAsset(repo)
	if err != nil {
		charmlog.Error("Failed to fetch specified repository", "repo", repo, "error", err)
		return nil, nil
	}
	// Fetch latest repo release
	l, err := release.FetchLatestRelease(repo)
	if err != nil {
		charmlog.Error("Failed to fetch latest release", "repo", repo, "error", err)
		return nil, nil
	}
	return a, l
}

func UpdateToLatestAsset(repo string, force bool) {
	currentAsset, latestAsset, _, latestTag, status := CheckUpdate(repo)

	// Replace if new version found/or nothing changes
	newVersion := util.UpdateVersion(currentAsset.AssetName, string(latestTag))

	if status {
		charmlog.Info("Current Version", currentAsset.Tag, "New Version", latestAsset.TagName)
		installUpdatedAsset(repo, latestAsset, currentAsset, newVersion, force)
	}

}
