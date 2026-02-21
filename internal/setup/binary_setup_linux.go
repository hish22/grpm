package setup

import (
	"hish22/grpm/internal/asset"
	"hish22/grpm/internal/link"
	"hish22/grpm/internal/util"
	"os"
	"path/filepath"

	charmlog "github.com/charmbracelet/log"
)

func MoveBinary(repo string, location string, assetID int, force bool) {
	if util.IsBinary(location) {
		binaryName := filepath.Base(location)
		// Make parent dir
		parentPath := link.WriteLibFilePath(repo)
		if err := os.MkdirAll(parentPath, 0755); err != nil {
			charmlog.Error("Failed to create parent directory", "error", err)
			return
		}
		newLink := filepath.Join(parentPath, binaryName)
		if err := os.Rename(location, newLink); err != nil {
			charmlog.Error("Failed to move binary from Downloads to lib", "error", err)
			return
		}
		SymlinkAsset(binaryName, newLink, binaryName, assetID, force)
		asset.InsertFileSetupLocation(newLink, assetID)
	}
}
