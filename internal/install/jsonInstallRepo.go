package install

import (
	"fmt"
	"hish22/grpm/internal/config"
	"hish22/grpm/internal/persistance"
	"hish22/grpm/internal/release"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3" // Import for side-effects

	"github.com/dustin/go-humanize"
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

func printAssets(r *release.Release, repo *string) {
	fmt.Println("=== Which asset of (", *repo, r.TagName, ") you want to install? ===")
	for i, a := range r.Assets {
		fmt.Print(i, "-")
		fmt.Println(a.AssetName, "("+humanize.Bytes(uint64(a.Size))+")")
	}
}

func FetchAssets(repo *string, tag *string) ([]release.Assets, *release.Release) {
	r := release.FetchSpecificRelease(repo, tag)
	printAssets(r, repo)
	return r.Assets, r
}

func FetchLatestReleaseAssets(repo *string) ([]release.Assets, *release.Release) {
	r := release.FetchLatestRelease(repo)
	printAssets(r, repo)
	return r.Assets, r
}

func writePath(assetName *string) string {
	configs := config.DecodeTOMLConfig()
	homePath, err := os.UserHomeDir()

	if err != nil {
		log.Fatal("Can't return home dir path")
	}

	return filepath.Join(homePath, configs.Downloaded, *assetName)
}

func InstallSelectedAsset(asset *release.Assets, release *release.Release) {
	// Fetch assets data and buffer it
	fmt.Printf("(%s) Installing..\n", asset.AssetName)
	resp, err := http.Get(asset.DownloadUrl)
	if err != nil {
		log.Fatal("Can't downloaded requested asset, ", asset.AssetName, err)
	}
	// Read assets data
	buf, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Can't read downloaded asset buffer, ", err)
	}
	// Create the file into the (Downloaded config location)
	path := writePath(&asset.AssetName)
	if err := os.WriteFile(path, buf, 0755); err != nil {
		log.Fatal("Can't write buffer data into a file,", err)
	}
	fmt.Printf("Asset %s installed successfully\n", asset.AssetName)

	trackAssetTable()             // Create the table if not exists
	registerAsset(asset, release) // Register installed asset

}

func trackAssetTable() {
	db := persistance.OpenMetadataDB()
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS asset (id INT, asset_name TEXT, location TEXT, tag TEXT, release_name TEXT, size INT, Digest TEXT);")
	if err != nil {
		log.Fatal("Can't create asset table to track assets, ", err)
	}
}

func registerAsset(asset *release.Assets, release *release.Release) {
	db := persistance.OpenMetadataDB()
	path := writePath(&asset.AssetName)
	_, err := db.Exec("INSERT INTO asset VALUES (?,?,?,?,?,?,?);", asset.ID, asset.AssetName, path, release.TagName, release.ReleaseName, asset.Size, asset.Digest)
	if err != nil {
		log.Fatal("Can't register an installed asset")
	}
	fmt.Println("Asset registered (Tracked)")
}
