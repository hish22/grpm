package persistance

import (
	"os"

	charmlog "github.com/charmbracelet/log"
)

func ClearCache() {
	// Remove grpm cache file
	root := ChacheRootLocation("grpm")
	err := os.RemoveAll(root)
	if err != nil {
		charmlog.Fatal("Failed to remove grpm cache file", "error", err)
	}

	// remove cache table from metadata.db
	db := OpenMetadataDB()
	_, err = db.Exec("DELETE FROM cache;")
	if err != nil {
		charmlog.Fatal("Failed to remove cache table from db", "error", err)
	}
}
