package list

import (
	"fmt"
	"hish22/grpm/internal/asset"

	charmlog "github.com/charmbracelet/log"
	"github.com/dustin/go-humanize"
)

func ListAssets() {
	trackedAssets, err := asset.FetchAssetsWithoutPrint()
	if err != nil {
		if err.Error() == "no such table: asset" {
			charmlog.Fatal("No installed assets to list,", "Error", err)
		} else {
			charmlog.Fatal(err)
		}
	}
	fmt.Println("=== Installed Assets (Tracked) ===")
	for _, a := range trackedAssets {
		fmt.Println(a.ID, a.AssetName, a.Tag, "("+humanize.Bytes(uint64(a.Size))+")")
	}
}
