package config

import (
	"errors"
	"os"
	"os/exec"
	"runtime"

	charmlog "github.com/charmbracelet/log"
)

func OpenTOMLConfig() error {
	switch runtime.GOOS {
	case "linux":
		cmd := exec.Command("nano", LocalConfigDirToml().String())

		// Connect the command's input/output to the current terminal
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err := cmd.Run()

		if err != nil {
			charmlog.Error("Failed to open TOML config file, ", "error", err)
		}

		return err

	case "windows":
		cmd := exec.Command("notepad", LocalConfigDirToml().String())
		err := cmd.Run()
		if err != nil {
			charmlog.Error("Failed to open TOML config file, ", "error", err)
		}
		return err
	}
	return errors.New("This command only supports (linux nano text editor) and (windows notepad text editor)")
}
