package asset

import (
	"hish22/grpm/internal/link"
	"hish22/grpm/internal/persistance"
	"hish22/grpm/internal/structures"

	charmlog "github.com/charmbracelet/log"
)

func TrackAssetTable() error {
	db := persistance.OpenMetadataDB()
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS asset (id INT PRIMARY KEY, repo TEXT,asset_name TEXT, location TEXT, tag TEXT, release_name TEXT, size INT, Digest TEXT, setup_track BOOL);")
	defer db.Close()
	if err != nil {
		charmlog.Error("Failed to create asset table to track assets, ", err)
	}
	return err
}

func RegisterAsset(repo string, asset *structures.Assets, release *structures.Release, setupTrack bool) {
	db := persistance.OpenMetadataDB()
	defer db.Close()
	path := link.WriteDownloadsFilePath(asset.AssetName)
	_, err := db.Exec("INSERT INTO asset VALUES (?,?,?,?,?,?,?,?,?);", asset.ID, repo, asset.AssetName, path, release.TagName, release.ReleaseName, asset.Size, asset.Digest, setupTrack)
	if err != nil {
		charmlog.Warn("Failed to register an installed asset", "error", err)
		return
	}
	charmlog.Info("Asset registered (Tracked)")
}

func AssetSetupTrackStatus(id int) bool {
	var status bool
	db := persistance.OpenMetadataDB()
	defer db.Close()
	row := db.QueryRow("SELECT setup_track FROM asset WHERE id=?", id)
	if row.Err() == nil {
		err := row.Scan(&status)
		if err != nil {
			charmlog.Error("Failed to fetch status of an asset", "error", err)
			return false
		}
	}
	return status
}
