package config

import (
	"fmt"
	"hish22/grpm/internal/util"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

func LocalConfigDirPath() string {
	return filepath.Join(util.HomeDir(), ".config", "grpm")
}

func LocalConfigDirToml() string {
	return filepath.Join(util.HomeDir(), ".config", "grpm", "config.toml")
}

func GenerateTOMLConfig() {

	payload := []byte(`location = "/usr/local/bin"` + "\n" +
		`downloaded = "/Downloads"` + "\n" +
		`arch = "` + runtime.GOARCH + `"` + "\n" +
		`os = "` + runtime.GOOS + `"`)

	if err := os.MkdirAll(LocalConfigDirPath(), 0755); err != nil {
		log.Fatal("Can't create grpm dir, ", err)
	}

	if err := os.WriteFile(LocalConfigDirToml(), payload, 0644); err != nil {
		log.Fatal("Can't create config.toml, ", err)
	}

	fmt.Printf("Config created at: %s\n", LocalConfigDirToml())

}
