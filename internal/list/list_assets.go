package list

import (
	"fmt"
	"hish22/grpm/internal/asset"

	"github.com/dustin/go-humanize"
)

func ListAssets() {
	ta := asset.FetchAssetsWithoutPrint()
	fmt.Println("=== Installed Assets (Tracked) ===")
	for _, a := range ta {
		fmt.Println(a.ID, a.AssetName, a.Tag, "("+humanize.Bytes(uint64(a.Size))+")")
	}
}
