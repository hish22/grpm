package config

import (
	"fmt"
	"hish22/grpm/internal/util"
	"os"
	"path/filepath"
	"runtime"

	charmlog "github.com/charmbracelet/log"
)

func LocalConfigDirPath() string {
	return filepath.Join(util.HomeDir(), ".config", "grpm")
}

func LocalConfigDirToml() string {
	return filepath.Join(util.HomeDir(), ".config", "grpm", "config.toml")
}

func GrpmDirPath() (string, error) {
	switch runtime.GOOS {
	case "linux":
		return filepath.Join("/", "opt", "grpm"), nil
	case "windows":
		return filepath.Join("C:\\", "Tools", "grpm"), nil
	default:
		return "", fmt.Errorf("Failed to retun the grpm path based on specified OS")
	}
}

func GrpmLibraryDirPath() (string, error) {
	switch runtime.GOOS {
	case "linux":
		return filepath.Join("/", "opt", "grpm", "lib"), nil
	case "windows":
		return filepath.Join("C:\\", "Tools", "grpm", "lib"), nil
	default:
		return "", fmt.Errorf("Failed to retun the grpm/lib path based on specified OS")
	}
}

func GrpmDownloadedDirPath() (string, error) {
	switch runtime.GOOS {
	case "linux":
		return filepath.Join("/", "opt", "grpm", "Downloads"), nil
	case "windows":
		return filepath.Join("C:\\", "Tools", "grpm", "Downloads"), nil
	default:
		return "", fmt.Errorf("Failed to retun the grpm/lib path based on specified OS")
	}
}

func createAndWriteConfig(payload []byte) {
	if err := os.MkdirAll(LocalConfigDirPath(), 0755); err != nil {
		charmlog.Fatal("Failed create .config/grpm dir, ", "error", err)
	}

	if err := os.WriteFile(LocalConfigDirToml(), payload, 0644); err != nil {
		charmlog.Fatal("Failed create config.toml, ", "error", err)
	}
}

func createGrpmDir() {

	gdd, err := GrpmDownloadedDirPath()

	if err != nil {
		charmlog.Fatal("Failed to fetch grpm and grpm/Downloads dir", "error", err)
	}

	libd, err := GrpmLibraryDirPath()

	if err != nil {
		charmlog.Fatal("Failed to fetch grpm/lib dir", "error", err)
	}

	if err := os.MkdirAll(gdd, 0755); err != nil {
		charmlog.Fatal("Failed to create grpm dir and grpm/Downloads dir, ", "error", err)
	}
	if err := os.MkdirAll(libd, 0755); err != nil {
		charmlog.Fatal("Failed to create grpm/lib dir", "error", err)
	}
}

func GenerateTOMLConfig() {
	var payload []byte
	switch runtime.GOOS {
	case "linux":
		payload = []byte(` location = "/opt/grpm"
			 library = "/opt/grpm/lib"` + "\n" +
			`downloaded = "/opt/grpm/Downloads"` + "\n" +
			`arch = "` + runtime.GOARCH + `"` + "\n" +
			`os = "` + runtime.GOOS + `"`)
	case "windows":
		payload = []byte(` location = "/Tools/grpm"
			 library = "/Tools/grpm/lib"` + "\n" +
			`downloaded = "/Tools/grpm/Downloads"` + "\n" +
			`arch = "` + runtime.GOARCH + `"` + "\n" +
			`os = "` + runtime.GOOS + `"`)
	}

	createAndWriteConfig(payload)
	createGrpmDir()
	fmt.Println(GrpmDirPath())
	charmlog.Info("Configuration initialized", "Location", LocalConfigDirToml())

}
