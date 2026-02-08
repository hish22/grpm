package install

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"hish22/grpm/internal/config"
	"hish22/grpm/internal/persistance"
	"hish22/grpm/internal/structures"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	charmlog "github.com/charmbracelet/log"
	_ "github.com/mattn/go-sqlite3" // Import for side-effects
)

// Installation of a pure releases (Without setup phase):
// (Default)
// 1- request/fetch a release
// 2- Display assets and select an asset to download
// 3- download the specified release
// 4- after the download completed register the asset into db
// 5- next time user asks to download same asset check db before downloading

// Installation with setup:
// 1- after the download step completed, ask user if he/she wants to setup the installed file
// 2- we handle the setup phase based on the file (type/extension)
// 3- zip/rar/7z (compressed files) have there own logic
// 4- reguler apps such cli tool also have there own logic
// 5- we need also to track the files

// Installation based on readme file:
// 1- read the readme file
// 2- extract most matched patterns (we could use agent)
// 3- apply steps
// 4- traksing could be opitional

func writePath(fileName *string) string {
	configs := config.DecodeTOMLConfig()
	homePath, err := os.UserHomeDir()

	if err != nil {
		charmlog.Fatal("Can't return home dir path")
	}

	return filepath.Join(homePath, configs.Downloaded, *fileName)
}

func downloadWithValidation(asset *structures.Assets, resp *http.Response) error {
	// Create .tmp file into specified download folder
	fileName := strconv.Itoa(asset.ID)
	tf, _ := os.CreateTemp("", fileName)
	defer os.Remove(tf.Name())
	defer tf.Close()

	// Create Hasher
	hash256 := sha256.New()
	// multiple stream write to tmp file and hasher
	mw := io.MultiWriter(tf, hash256)
	io.Copy(mw, resp.Body)

	// compare digest
	calcDigest := "sha256:" + hex.EncodeToString(hash256.Sum(nil))
	if asset.Digest != calcDigest {
		return fmt.Errorf("Digest Unmatch!")
	}

	tf.Close()
	return os.Rename(tf.Name(), writePath(&asset.AssetName))
}

func InstallSelectedAsset(repo *string, asset *structures.Assets, release *structures.Release) {
	// Request to Fetch assets from specific release
	charmlog.Info("Installing..", "asset", asset.AssetName)
	resp, err := http.Get(asset.DownloadUrl)
	if err != nil {
		charmlog.Fatal("Failed to GET request asset payload, ", "asset", asset.AssetName, "error", err)
	}

	// Create .tmp file where we store binary data
	// to validate the digest of the fetched content
	if err := downloadWithValidation(asset, resp); err != nil {
		charmlog.Fatal("Failed to download asset", "asset", asset.AssetName, "error", err)
	}

	charmlog.Info("Digest match")

	trackAssetTable()                   // Create the table if not exists
	registerAsset(repo, asset, release) // Register installed asset

}

func trackAssetTable() {
	db := persistance.OpenMetadataDB()
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS asset (id INT, repo TEXT,asset_name TEXT, location TEXT, tag TEXT, release_name TEXT, size INT, Digest TEXT);")
	if err != nil {
		charmlog.Fatal("Can't create asset table to track assets, ", err)
	}
}

func registerAsset(repo *string, asset *structures.Assets, release *structures.Release) {
	db := persistance.OpenMetadataDB()
	path := writePath(&asset.AssetName)
	_, err := db.Exec("INSERT INTO asset VALUES (?,?,?,?,?,?,?,?);", asset.ID, *repo, asset.AssetName, path, release.TagName, release.ReleaseName, asset.Size, asset.Digest)
	if err != nil {
		charmlog.Fatal("Can't register an installed asset")
	}
	charmlog.Info("Asset registered (Tracked)")
}
