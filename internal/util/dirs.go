package util

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

func HomeDir() (string, error) {
	switch runtime.GOOS {
	case "linux":
		if os.Geteuid() == 0 {
			return filepath.Join("/", "home", os.Getenv("SUDO_USER")), nil
		} else {
			home, err := os.UserHomeDir()
			if err != nil {
				return "", err
			}
			return home, nil
		}
	case "windows":
		home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		return home, nil
	default:
		return "", fmt.Errorf("Failed to fetch home path")
	}
}
