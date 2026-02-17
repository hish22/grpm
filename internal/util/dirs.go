package util

import (
	"os"
	"os/user"
)

func HomeDir() (string, error) {
	sudoUser := os.Getenv("SUDO_USER")
	if sudoUser != "" {
		// We are in sudo, find the original user's home
		u, err := user.Lookup(sudoUser)
		if err != nil {
			return "", err
		}
		return u.HomeDir, nil
	}
	// Not in sudo, use standard Go function
	return os.UserHomeDir()

	// switch runtime.GOOS {
	// case "linux":
	// 	if os.Geteuid() == 0 {
	// 		return filepath.Join("/", "home", os.Getenv("SUDO_USER")), nil
	// 	} else {
	// 		home, err := os.UserHomeDir()
	// 		if err != nil {
	// 			return "", err
	// 		}
	// 		return home, nil
	// 	}
	// case "windows":
	// 	home, err := os.UserHomeDir()
	// 	if err != nil {
	// 		return "", err
	// 	}
	// 	return home, nil
	// default:
	// 	return "", fmt.Errorf("Failed to fetch home path")
	// }
}
