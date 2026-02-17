package list

import (
	"fmt"
	"hish22/grpm/internal/asset"
	"strconv"

	"github.com/charmbracelet/lipgloss"
	charmlog "github.com/charmbracelet/log"
	"github.com/dustin/go-humanize"
)

var (
	idstyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("#2b77fb")).Bold(true)
	tagstyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#ff0000"))
)

func ListAssets() {
	trackedAssets, err := asset.FetchAssets()
	if err != nil {
		if err.Error() == "no such table: asset" {
			charmlog.Fatal("No installed assets to list,", "Error", err)
		} else {
			charmlog.Fatal(err)
		}
	}
	if len(trackedAssets) <= 0 {
		fmt.Println("No installed assets.")
	} else {
		fmt.Println("=== Installed Assets (Tracked) ===")
		for _, a := range trackedAssets {
			fmt.Println(idstyle.Render(strconv.Itoa(a.ID)), "["+a.RepoName+"]", a.AssetName, tagstyle.Render(a.Tag),
				"("+humanize.Bytes(uint64(a.Size))+")")
		}
	}

}
