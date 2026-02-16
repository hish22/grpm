package config

import (
	"fmt"
	"hish22/grpm/internal/util"
	"os"
	"path/filepath"
	"runtime"

	charmlog "github.com/charmbracelet/log"
)

type FileLink struct {
	Base     string
	Childern []string
	Asset    string
}

func (link FileLink) String() string {
	childern := filepath.Join(link.Childern...)
	return filepath.Join(link.Base, childern, link.Asset)
}

func LocalConfigDirPath() FileLink {
	home, err := util.HomeDir()
	if err != nil {
		charmlog.Fatal("Failed to fetch local config dirctory path", "error", err)
	}
	return FileLink{
		Base:     home,
		Childern: []string{".config", "grpm"},
	}
}

func LocalConfigDirToml() FileLink {
	home, err := util.HomeDir()
	if err != nil {
		charmlog.Fatal("Failed to fetch local config .toml", "error", err)
	}
	return FileLink{
		Base:     home,
		Childern: []string{".config", "grpm"},
		Asset:    "config.toml",
	}
}

func GrpmDirPath() (FileLink, error) {
	switch runtime.GOOS {
	case "linux":
		return FileLink{Base: "/", Childern: []string{"opt", "grpm"}}, nil
	case "windows":
		return FileLink{Base: "C:\\", Childern: []string{"Tools", "grpm"}}, nil
	default:
		return FileLink{}, fmt.Errorf("Failed to retun the grpm path based on specified OS")
	}
}

func GrpmLibraryDirPath() (FileLink, error) {
	switch runtime.GOOS {
	case "linux":
		return FileLink{Base: "/", Childern: []string{"opt", "grpm", "lib"}}, nil
	case "windows":
		return FileLink{Base: "C:\\", Childern: []string{"Tools", "grpm", "lib"}}, nil
	default:
		return FileLink{}, fmt.Errorf("Failed to retun the grpm/lib path based on specified OS")
	}
}

func GrpmDownloadedDirPath() (FileLink, error) {
	switch runtime.GOOS {
	case "linux":
		return FileLink{Base: "/", Childern: []string{"opt", "grpm", "Downloads"}}, nil
	case "windows":
		return FileLink{Base: "C:\\", Childern: []string{"Tools", "grpm", "Downloads"}}, nil
	default:
		return FileLink{}, fmt.Errorf("Failed to retun the grpm/lib path based on specified OS")
	}
}

func createAndWriteConfig(payload []byte) {
	if err := os.MkdirAll(LocalConfigDirPath().String(), 0755); err != nil {
		charmlog.Fatal("Failed create .config/grpm dir, ", "error", err)
	}

	if err := os.WriteFile(LocalConfigDirToml().String(), payload, 0644); err != nil {
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

	if err := os.MkdirAll(gdd.String(), 0755); err != nil {
		charmlog.Fatal("Failed to create grpm dir and grpm/Downloads dir, ", "error", err)
	}
	if err := os.MkdirAll(libd.String(), 0755); err != nil {
		charmlog.Fatal("Failed to create grpm/lib dir", "error", err)
	}
}

func GenerateTOMLConfig() {
	var payload []byte
	switch runtime.GOOS {
	case "linux":
		payload = []byte(`location = "/opt/grpm"` + "\n" +
			`library = "/opt/grpm/lib"` + "\n" +
			`downloaded = "/opt/grpm/Downloads"` + "\n" +
			`arch = "` + runtime.GOARCH + `"` + "\n" +
			`os = "` + runtime.GOOS + `"`)
	case "windows":
		payload = []byte(`location = "/Tools/grpm"` + "\n" +
			`library = "/Tools/grpm/lib"` + "\n" +
			`downloaded = "/Tools/grpm/Downloads"` + "\n" +
			`arch = "` + runtime.GOARCH + `"` + "\n" +
			`os = "` + runtime.GOOS + `"`)
	}
	createAndWriteConfig(payload)
	createGrpmDir()
	charmlog.Info("Configuration initialized", "Location", LocalConfigDirToml().String())
}
