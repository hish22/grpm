package config

import (
	"os"
)

func CheckConfig() bool {
	_, err := os.Stat(LocalConfigDirToml())
	if err != nil {
		return false
	}
	return true
}
