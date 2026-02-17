package asset

import (
	"fmt"
	"hish22/grpm/internal/persistance"
	"os"

	charmlog "github.com/charmbracelet/log"
)

func DeleteLastTrackedAssetById(id int) {
	db := persistance.OpenMetadataDB()
	_, err := db.Exec("DELETE FROM asset WHERE id=?", id)
	if err != nil {
		charmlog.Error("Failed to delete last tracked asset", "error", err)
	}
}

func RemoveRawAsset(assetLocation string) error {
	err := os.RemoveAll(assetLocation)
	if err != nil {
		charmlog.Error("Failed to delete old raw asset folder", "error", err)
		return err
	}
	return nil
}

func RemoveAssetLibFile(assetID int) error {
	assetLibLocation := FileSetupLocation(assetID)
	if assetLibLocation != "" {
		err := os.RemoveAll(assetLibLocation)
		if err != nil {
			charmlog.Error("Failed to delete old asset folder", "error", err)
			return err
		}
	} else {
		return fmt.Errorf("No such setup location for intended asset")
	}
	return nil
}
