package install

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	assets "hish22/grpm/internal/asset"
	"hish22/grpm/internal/link"
	"hish22/grpm/internal/setup"
	"hish22/grpm/internal/structures"
	"hish22/grpm/internal/util"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

	charmlog "github.com/charmbracelet/log"
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

func isInstalled(repo string) (*structures.TrackedAsset, bool) {
	asset, err := assets.FetchSpecificAsset(repo)
	if err != nil {
		return &asset, false
	}
	return &asset, true
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

	charmlog.Info("Comparing digests..")
	// compare digest
	calcDigest := "sha256:" + hex.EncodeToString(hash256.Sum(nil))
	if asset.Digest != calcDigest {
		return fmt.Errorf("Digest Unmatch!")
	}
	tf.Close()
	return os.Rename(tf.Name(), link.WriteDownloadsFilePath(asset.AssetName))
}

func changeFilePerm(asset string) {
	err := os.Chmod(link.WriteDownloadsFilePath(asset), 0644)
	if err != nil {
		charmlog.Fatal("Failed ot change file's permission", "error", err)
	}
}

func InstallSelectedAsset(repo string, asset *structures.Assets, release *structures.Release, setupStatus bool) {
	// Check if asset is already installed
	installedAsset, isInstalled := isInstalled(repo)
	if isInstalled {
		charmlog.Info("Asset is already installed", "asset", installedAsset.Location)
		return
	}

	// Request to Fetch assets from specific release
	charmlog.Info("Installing..", "asset", asset.AssetName)
	resp, err := http.Get(asset.DownloadUrl)
	if err != nil {
		charmlog.Fatal("Failed to GET request asset payload", "asset", asset.AssetName, "error", err)
	}

	// Create .tmp file where we store binary data
	// to validate the digest of the fetched content
	// and handle unexpected situations
	if err := downloadWithValidation(asset, resp); err != nil {
		charmlog.Fatal("Failed to download asset (Make sure to use sudo/privileged execution)", "asset", asset.AssetName, "error", err)
	}

	charmlog.Info("Digest match")

	changeFilePerm(asset.AssetName)

	assets.TrackAssetTable()                                // Create the table if not exists
	assets.RegisterAsset(repo, asset, release, setupStatus) // Register installed asset

	// auto setup of installed file
	// if the user only flaged with --setup
	if setupStatus {
		exts := util.ExtensionExtractor(asset.AssetName)
		totalext := strings.Join(exts, "")
		setup.SetupAsset(link.WriteDownloadsFilePath(asset.AssetName), totalext, asset.ID)
		charmlog.Info("Asset installed at /opt/grpm/lib")
	}
}
