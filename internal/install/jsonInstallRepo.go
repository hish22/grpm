package install

import (
	"fmt"
	"hish22/grpm/internal/config"
	"hish22/grpm/internal/release"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/dustin/go-humanize"
)

// Default
// 1- request/fetch a release
// 2- Display assets and select an asset to download
// 3- after the download step completed, ask user if he/she wants to setup the installed file

func printAssets(r *release.Release) {
	fmt.Println("=== Which asset of (", (*r).ReleaseName, ") you want to install? ===")
	for i, a := range r.Assets {
		fmt.Print(i, "-")
		fmt.Println(a.AssetName, "("+humanize.Bytes(uint64(a.Size))+")")
	}
}

func FetchAssets(repo *string, tag *string) []release.Assets {
	r := release.FetchSpecificRelease(repo, tag)
	printAssets(r)
	return r.Assets
}

func FetchLatestReleaseAssets(repo *string) []release.Assets {
	r := release.FetchLatestRelease(repo)
	printAssets(r)
	return r.Assets
}

func InstallSelectedAsset(asset *release.Assets) {
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
	configs := config.DecodeTOMLConfig()
	homePath, err := os.UserHomeDir()

	if err != nil {
		log.Fatal("Can't return home dir path")
	}

	path := filepath.Join(homePath, configs.Downloaded, asset.AssetName)
	if err := os.WriteFile(path, buf, 0755); err != nil {
		log.Fatal("Can't write buffer data into a file,", err)
	}
	fmt.Printf("Asset %s installed successfully\n", asset.AssetName)
}
