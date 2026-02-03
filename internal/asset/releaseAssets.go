package asset

import (
	"fmt"
	"hish22/grpm/internal/config"
	"hish22/grpm/internal/persistance"
	"hish22/grpm/internal/release"
	"log"
	"strings"

	"github.com/dustin/go-humanize"
)

func printAssets(r []release.Assets) {
	for i, a := range r {
		fmt.Print(i, "-")
		fmt.Println(a.AssetName, "("+humanize.Bytes(uint64(a.Size))+")")
	}
}

func matchedAssets(r *release.Release) []release.Assets {
	config := config.DecodeTOMLConfig()
	var MatchedReleases []release.Assets
	for _, a := range r.Assets {
		if strings.Contains(a.AssetName, config.Arch) && strings.Contains(a.AssetName, config.Os) {
			MatchedReleases = append(MatchedReleases, a)
		}
	}
	return MatchedReleases
}

func PrintTheAssets(r *release.Release, repo *string, match bool) {
	fmt.Println("=== Which asset of (", *repo, r.TagName, ") you want to install? ===")
	if match {
		a := matchedAssets(r)
		printAssets(a)
	} else {
		printAssets(r.Assets)
	}
}

func FetchAssets(repo *string, tag *string) ([]release.Assets, *release.Release) {
	r := release.FetchSpecificRelease(repo, tag)
	return r.Assets, r
}

func FetchLatestReleaseAssets(repo *string) ([]release.Assets, *release.Release) {
	r := release.FetchLatestRelease(repo)
	return r.Assets, r
}

func FetchAssetsWithoutPrint() []release.Assets {
	db := persistance.OpenMetadataDB()
	assets := []release.Assets{}
	r, err := db.Query("SELECT * FROM asset")
	if err != nil {
		log.Fatal("Can't fetch installed assets")
	}
	defer r.Close()
	if r.Next() {
		r.Scan(assets)
	}
	return assets
}
