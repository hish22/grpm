package asset

import (
	"hish22/grpm/internal/persistance"

	charmlog "github.com/charmbracelet/log"
)

func DeleteLastTrackedAssetById(id int) {
	db := persistance.OpenMetadataDB()
	_, err := db.Exec("DELETE FROM asset WHERE id=?", id)
	if err != nil {
		charmlog.Fatal("Failed to delete last tracked asset", "error", err)
	}
}
