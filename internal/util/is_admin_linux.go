package util

import "os"

func IsAdministrator() bool {
	if os.Getuid() == 0 {
		return true
	} else {
		return false
	}
}
