package asset

import (
	"context"
	"hish22/grpm/internal/link"
	"hish22/grpm/internal/middlewares"
	"hish22/grpm/internal/structures"
	"time"

	charmlog "github.com/charmbracelet/log"
)

func TrackAssetTable() error {
	db := middlewares.MetadataDBConn()
	ctx, cancle := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancle()
	_, err := db.ExecContext(ctx, "CREATE TABLE IF NOT EXISTS asset (id INT PRIMARY KEY, repo TEXT,asset_name TEXT, location TEXT, tag TEXT, release_name TEXT, size INT, Digest TEXT, setup_track BOOL, symlink_orenv_name TEXT,file_setup_location TEXT);")
	defer db.Close()
	if err != nil {
		charmlog.Error("Failed to create asset table to track assets, ", err)
	}
	return err
}

func RegisterAsset(repo string, asset *structures.Assets, release *structures.Release, setupTrack bool) {
	db := middlewares.MetadataDBConn()
	defer db.Close()
	path := link.WriteDownloadsFilePath(asset.AssetName)
	ctx, cancle := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancle()
	_, err := db.ExecContext(ctx, "INSERT INTO asset VALUES (?,?,?,?,?,?,?,?,?,?,?);", asset.ID, repo, asset.AssetName, path, release.TagName, release.ReleaseName, asset.Size, asset.Digest, setupTrack, "", "")
	if err != nil {
		charmlog.Warn("Failed to register an installed asset", "error", err)
		return
	}
	charmlog.Info("Asset registered (Tracked)")
}

func InsertFileSetupLocation(location string, id int) {
	db := middlewares.MetadataDBConn()
	defer db.Close()
	ctx, cancle := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancle()
	_, err := db.ExecContext(ctx, "UPDATE asset SET file_setup_location=? WHERE id=?", location, id)
	if err != nil {
		charmlog.Error("Failed to insert file_setup_location to an asset", "error", err)
		return
	}
	charmlog.Info("File Setup location inserted")
}

func FileSetupLocation(id int) string {
	var location string
	db := middlewares.MetadataDBConn()
	defer db.Close()
	ctx, cancle := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancle()
	row := db.QueryRowContext(ctx, "SELECT file_setup_location FROM asset WHERE id=?", id)
	err := row.Scan(&location)
	if err != nil {
		charmlog.Error("Failed to fetch file_setup_location", "error", err)
	}
	return location
}

func AssetSetupTrackStatus(id int) bool {
	var status bool
	db := middlewares.MetadataDBConn()
	defer db.Close()
	ctx, cancle := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancle()
	row := db.QueryRowContext(ctx, "SELECT setup_track FROM asset WHERE id=?", id)
	if row.Err() == nil {
		err := row.Scan(&status)
		if err != nil {
			charmlog.Error("Failed to fetch status of an asset", "error", err)
			return false
		}
	}
	return status
}

func InsertSymlinkOrEnvLocation(name string, id int) {
	db := middlewares.MetadataDBConn()
	defer db.Close()
	ctx, cancle := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancle()
	_, err := db.ExecContext(ctx, "UPDATE asset SET symlink_orenv_name=? WHERE id=?", name, id)
	if err != nil {
		charmlog.Error("Failed to insert symlink_name to an asset", "error", err)
		return
	}
	charmlog.Info("Symlink or Environment variable location inserted")
}

func SymlinkOrEnvLocation(id int) string {
	var symlink string
	db := middlewares.MetadataDBConn()
	defer db.Close()
	ctx, cancle := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancle()
	row := db.QueryRowContext(ctx, "SELECT symlink_orenv_name FROM asset WHERE id=?", id)
	err := row.Scan(&symlink)
	if err != nil {
		charmlog.Error("Failed to fetch symlink_name", "error", err)
	}
	return symlink
}
