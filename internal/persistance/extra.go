package persistance

import (
	"context"
	"hish22/grpm/internal/middlewares"
	"os"
	"time"

	charmlog "github.com/charmbracelet/log"
)

func ClearCache() {
	// Remove grpm cache file
	root := ChacheRootLocation("grpm")
	err := os.RemoveAll(root)
	if err != nil {
		charmlog.Error("Failed to remove grpm cache file", "error", err)
		return
	}

	// remove cache table from metadata.db
	db := middlewares.MetadataDBConn()
	ctx, cancle := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancle()
	_, err = db.ExecContext(ctx, "DELETE FROM cache;")
	if err != nil {
		charmlog.Error("Failed to remove cache table from db", "error", err)
	}
}
