package util

import (
	"log"
	"os"
)

func HomeDir() string {
	homep, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("Can't detect home dir, ", err)
	}
	return homep
}
