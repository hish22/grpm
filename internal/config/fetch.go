package config

import (
	"hish22/grpm/internal/util"
	"log"
	"os"
	"path/filepath"
)

func fetchTOMLconfig() string {
	homep, err := util.HomeDir()
	if err != nil {
		log.Fatal("Can't detect home dir, ", err)
	}

	localConfigDirToml := filepath.Join(homep, ".config", "grpm", "config.toml")

	data, err := os.ReadFile(localConfigDirToml)

	if err != nil {
		log.Fatal("Can't read config.toml file,", err)
	}

	return string(data)

}
