package config

import (
	"log"
	"os"
	"os/exec"
	"runtime"
)

func OpenTOMLConfig() {
	switch runtime.GOOS {
	case "linux":
		cmd := exec.Command("nano", LocalConfigDirToml().String())

		// Connect the command's input/output to the current terminal
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			log.Fatal("Can't open TOML config file, ", err)
		}
	case "windows":
		cmd := exec.Command("notepad", LocalConfigDirToml().String())
		if err := cmd.Run(); err != nil {
			log.Fatal("Can't open TOML config file, ", err)
		}
	}
}
