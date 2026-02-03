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

type TrackedAsset struct {
	ID          int
	AssetName   string
	Location    string
	Tag         string
	ReleaseName string
	Size        int
	Digest      string
}

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
		r.Assets = matchedAssets(r)
	}
	printAssets(r.Assets)
}

func FetchAssetsWithoutPrint() []TrackedAsset {
	db := persistance.OpenMetadataDB()
	var a TrackedAsset
	assets := []TrackedAsset{}
	r, err := db.Query("SELECT * FROM asset;")
	if err != nil {
		log.Fatal("Can't fetch installed assets")
	}
	defer r.Close()
	if r.Next() {
		err := r.Scan(&a.ID, &a.AssetName, &a.Location, &a.Tag,
			&a.ReleaseName, &a.Size, &a.Digest)
		if err != nil {
			log.Fatal("Can't decode sql, ", err)
		}
		assets = append(assets, a)
	}
	return assets
}
