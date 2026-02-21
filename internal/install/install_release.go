package install

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	assets "hish22/grpm/internal/asset"
	"hish22/grpm/internal/config"
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
	"github.com/schollz/progressbar/v3"
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

/* Hit DB to check if a sepcfifc asset is already installed */
func isInstalled(repo string) (*structures.TrackedAsset, bool) {
	asset, err := assets.FetchSpecificAsset(repo)
	if err != nil {
		return &asset, false
	}
	return &asset, true
}

/*
	Start The Download operation

Fetch grpm/Downloads path first, then we insert the http response body payload
into a .tmp file. provide progress bar + hashing to check asset digest.
if digest check succeed, we start by renaming/creating the fetched asset.
*/
func downloadWithValidation(asset *structures.Assets, resp *http.Response) error {
	// Bring downloads (as tmp file) location
	downloads, err := config.GrpmDownloadedDirPath()
	if err != nil {
		return err
	}

	// Create .tmp file into specified download folder
	fileName := strconv.Itoa(asset.ID)
	tf, _ := os.CreateTemp(downloads.String(), fileName)
	defer os.Remove(tf.Name())
	defer tf.Close()

	// Progress bar
	bar := progressbar.DefaultBytes(
		int64(asset.Size),
		"Downloading",
	)
	// Create Hasher
	hash256 := sha256.New()
	// multiple stream write to tmp file and hasher
	mw := io.MultiWriter(tf, hash256, bar)
	io.Copy(mw, resp.Body)

	charmlog.Info("Comparing digests..")
	// compare digest
	calcDigest := "sha256:" + hex.EncodeToString(hash256.Sum(nil))
	if asset.Digest != calcDigest {
		return fmt.Errorf("Digest Unmatch!")
	} else {
		charmlog.Info("Digest match")
	}

	tf.Close()
	return os.Rename(tf.Name(), link.WriteDownloadsFilePath(asset.AssetName))
}

/*
	Change file permissions

To allow a user to extract and use the download file,
by changing the file's permission mode into 0644
*/
func changeFilePerm(asset string) {
	err := os.Chmod(link.WriteDownloadsFilePath(asset), 0644)
	if err != nil {
		charmlog.Fatal("Failed ot change file's permission", "error", err)
	}
}

/*
Install an asset process

Check if user is running this with privileged execution
then, Check if asset is already installed.

Request to Fetch assets from specific release.

start with setup process only if the user flaged with --setup.
*/
func InstallSelectedAsset(repo string, asset *structures.Assets, release *structures.Release, setupStatus bool, force bool) {
	// Check if user is running this with privileged execution
	if !util.IsAdministrator() {
		charmlog.Error("Please run this command with privilege execution mode")
		return
	}

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
	defer resp.Body.Close()
	// Create .tmp file where we store binary data
	// to validate the digest of the fetched content
	// and handle unexpected situations
	if err := downloadWithValidation(asset, resp); err != nil {
		charmlog.Fatal("Failed to download the asset", "asset", asset.AssetName, "error", err)
	}

	changeFilePerm(asset.AssetName)

	assets.TrackAssetTable()                                // Create the table if not exists
	assets.RegisterAsset(repo, asset, release, setupStatus) // Register installed asset

	// auto setup of installed file
	// if the user only flaged with --setup
	if setupStatus {
		exts := util.ExtensionExtractor(asset.AssetName)
		totalext := strings.Join(exts, "")
		setup.SetupAsset(util.RepoNameExtractor(repo), link.WriteDownloadsFilePath(asset.AssetName),
			totalext, asset.ID, force)
		libLink, err := config.GrpmLibraryDirPath()
		if err != nil {
			charmlog.Warn("Failed to fetch grpm library directory")
		}
		charmlog.Info("Asset installed", "location", libLink.String())
	}
}
