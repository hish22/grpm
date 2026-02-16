package config

import (
	"os"
)

func CheckConfig() bool {
	_, err := os.Stat(LocalConfigDirToml().String())
	if err != nil {
		return false
	}
	return true
}
