package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func GenerateTOMLConfig() {

	homep, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("Can't detect home dir, ", err)
	}

	localConfigDirPath := filepath.Join(homep, ".config", "grpm")
	localConfigDirToml := filepath.Join(homep, ".config", "grpm", "config.toml")

	payload := []byte(`location = "/usr/local/bin"`)

	if err := os.MkdirAll(localConfigDirPath, 0755); err != nil {
		log.Fatal("Can't create grpm dir, ", err)
	}

	if err := os.WriteFile(localConfigDirToml, payload, 0644); err != nil {
		log.Fatal("Can't create config.toml, ", err)
	}

	fmt.Printf("Config created at: %s\n", localConfigDirToml)

}
