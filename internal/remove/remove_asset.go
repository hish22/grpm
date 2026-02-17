package remove

import (
	"hish22/grpm/internal/asset"
	"hish22/grpm/internal/setup"

	charmlog "github.com/charmbracelet/log"
)

func RemoveAssetByRepoName(repo string) {
	trackedAsset, err := asset.FetchSpecificAsset(repo)
	if err != nil {
		charmlog.Error("Failed to remove asset (Are you sure you have installed this asset?)", "error", err)
		return
	}
	setup.RemoveSymlink(trackedAsset.ID)
	err = asset.RemoveRawAsset(trackedAsset.Location)
	if err != nil {
		charmlog.Error("Failed to remove raw asset", "error", err)
		return
	}
	err = asset.RemoveAssetLibFile(trackedAsset.ID)
	if err != nil {
		charmlog.Error("Failed to remove lib file", "error", err)
	}
	asset.DeleteLastTrackedAssetById(trackedAsset.ID)
	charmlog.Info("Asset removed", "repository", repo)
}

func RemoveAssetByID(id int, location string) {
	setup.RemoveSymlink(id)
	err := asset.RemoveRawAsset(location)
	if err != nil {
		charmlog.Error("Failed to remove raw asset", "error", err)
		return
	}
	err = asset.RemoveAssetLibFile(id)
	if err != nil {
		charmlog.Error("Failed to remove lib file", "error", err)
	}
	asset.DeleteLastTrackedAssetById(id)
	charmlog.Info("Asset removed", "id", id)
}
