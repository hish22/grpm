package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func homeDir() string {
	homep, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("Can't detect home dir, ", err)
	}
	return homep
}

func LocalConfigDirPath() string {
	return filepath.Join(homeDir(), ".config", "grpm")
}

func LocalConfigDirToml() string {
	return filepath.Join(homeDir(), ".config", "grpm", "config.toml")
}

func GenerateTOMLConfig() {

	payload := []byte(`location = "/usr/local/bin"`)

	if err := os.MkdirAll(LocalConfigDirPath(), 0755); err != nil {
		log.Fatal("Can't create grpm dir, ", err)
	}

	if err := os.WriteFile(LocalConfigDirToml(), payload, 0644); err != nil {
		log.Fatal("Can't create config.toml, ", err)
	}

	fmt.Printf("Config created at: %s\n", LocalConfigDirToml())

}
