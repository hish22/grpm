package asset

import (
	"context"
	"fmt"
	"hish22/grpm/internal/config"
	"hish22/grpm/internal/middlewares"
	"hish22/grpm/internal/structures"
	"hish22/grpm/internal/util"
	"log"
	"strings"
	"time"

	"github.com/dustin/go-humanize"
)

func enumerateAssets(r []structures.Assets) {
	for i, a := range r {
		fmt.Print(i, "-")
		fmt.Println(a.AssetName, "("+humanize.Bytes(uint64(a.Size))+")")
	}
}

func matchedAssets(r *structures.Release) []structures.Assets {
	config := config.DecodeTOMLConfig()
	var matchedAssets []structures.Assets
	for _, a := range r.Assets {
		if strings.Contains(strings.ToLower(a.AssetName), config.Arch) &&
			strings.Contains(strings.ToLower(a.AssetName), config.Os) {
			matchedAssets = append(matchedAssets, a)
		} else if strings.Contains(strings.ToLower(a.AssetName), config.Os) {
			util.ArchitectureAssetsMatch(&config.Arch, &a, &matchedAssets)
		}
	}
	return matchedAssets
}

func PrintTheAssets(r *structures.Release, repo string, match bool) {
	fmt.Println("=== Which asset of (", repo, r.TagName, ") you want to install? ===")
	if match {
		r.Assets = matchedAssets(r)
	}
	enumerateAssets(r.Assets)
}

func FetchAssets() ([]structures.TrackedAsset, error) {
	db := middlewares.MetadataDBConn()
	defer db.Close()
	var a structures.TrackedAsset
	assets := []structures.TrackedAsset{}
	ctx, cancle := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancle()
	r, err := db.QueryContext(ctx, "SELECT * FROM asset;")
	if err != nil {
		return nil, err
	}
	defer r.Close()
	for r.Next() {
		err := r.Scan(&a.ID, &a.RepoName, &a.AssetName, &a.Location, &a.Tag,
			&a.ReleaseName, &a.Size, &a.Digest, &a.SetupStatus, &a.SymlinkName, &a.FileSetupLocation)
		if err != nil {
			log.Fatal("Can't decode sql, ", err)
		}
		assets = append(assets, a)
	}
	return assets, nil
}

func FetchSpecificAsset(repo string) (*structures.TrackedAsset, error) {
	db := middlewares.MetadataDBConn()
	defer db.Close()
	ctx, cancle := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancle()
	row := db.QueryRowContext(ctx, "SELECT * FROM asset WHERE repo=?", repo)
	if row.Err() != nil {
		return &structures.TrackedAsset{}, row.Err()
	}
	a := structures.TrackedAsset{}
	err := row.Scan(&a.ID, &a.RepoName, &a.AssetName, &a.Location, &a.Tag, &a.ReleaseName,
		&a.Size, &a.Digest, &a.SetupStatus, &a.SymlinkName, &a.FileSetupLocation)
	if err != nil {
		return &structures.TrackedAsset{}, err
	}
	return &a, nil
}
