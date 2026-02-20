package setup

import (
	"bufio"
	"fmt"
	"hish22/grpm/internal/asset"
	"hish22/grpm/internal/config"
	"hish22/grpm/internal/util"
	"os"
	"strings"

	charmlog "github.com/charmbracelet/log"
)

func enableExecute(asset string) {
	charmlog.Info("Changing permissions")
	err := os.Chmod(asset, 0755)
	if err != nil {
		charmlog.Error("Failed to change permissions mode as 744", "error", err)
	}
}

func confirm(binaryName string, force bool) bool {
	charmlog.Info("Binary detected", "binary", binaryName)
	fmt.Print("Do you want to create a symlink for this binary (yes/no)? ")
	status := false

	// apply without asking to enter yes or no
	if force {
		return true
	}

	scan := bufio.NewScanner(os.Stdin)
	if scan.Scan() {
		switch scan.Text() {
		case "yes":
			status = true
		case "y":
			status = true
		case "no":
			status = false
		case "n":
			status = false
		}
	}
	return status
}

func SymlinkAsset(repo string, assetLocation string, binaryName string, assetID int, force bool) {
	if util.IsBinary(assetLocation) && (strings.EqualFold(repo, binaryName) || strings.Contains(binaryName, repo)) && confirm(binaryName, force) {
		enableExecute(assetLocation)
		newlink := config.FileLink{
			Base:     "/",
			Childern: []string{"usr", "local", "bin"},
			Asset:    binaryName,
		}
		err := os.Symlink(assetLocation, newlink.String())
		if err != nil {
			charmlog.Error("Failed to create symlink to binary", "binary", binaryName, "error", err)
		}
		charmlog.Info("Symlink created", "asset", binaryName, "location", newlink.String())
		asset.InsertSymlinkOrEnvLocation(binaryName, assetID)
	}
}

func RemoveSymlink(id int) {
	symlinkName := asset.SymlinkOrEnvLocation(id)
	if symlinkName != "" {
		link := config.FileLink{
			Base:     "/",
			Childern: []string{"usr", "local", "bin"},
			Asset:    symlinkName,
		}
		err := os.Remove(link.String())
		if err != nil {
			charmlog.Error("Failed to remove old symlink", "error", err)
		}
	} else {
		charmlog.Info("No symlink created for this asset")
	}
}
