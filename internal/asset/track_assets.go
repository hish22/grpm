package asset

import (
	corehttp "hish22/grpm/internal/coreHttp"
	"hish22/grpm/internal/persistance"
	"hish22/grpm/internal/structures"

	charmlog "github.com/charmbracelet/log"
)

func TrackAssetTable() {
	db := persistance.OpenMetadataDB()
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS asset (id INT PRIMARY KEY, repo TEXT,asset_name TEXT, location TEXT, tag TEXT, release_name TEXT, size INT, Digest TEXT);")
	if err != nil {
		charmlog.Fatal("Can't create asset table to track assets, ", err)
	}
}

func RegisterAsset(repo *string, asset *structures.Assets, release *structures.Release) {
	db := persistance.OpenMetadataDB()
	path := corehttp.WriteFilePath(&asset.AssetName)
	_, err := db.Exec("INSERT INTO asset VALUES (?,?,?,?,?,?,?,?);", asset.ID, *repo, asset.AssetName, path, release.TagName, release.ReleaseName, asset.Size, asset.Digest)
	if err != nil {
		charmlog.Warn("Failed to register an installed asset", "error", err)
		return
	}
	charmlog.Info("Asset registered (Tracked)")
}
