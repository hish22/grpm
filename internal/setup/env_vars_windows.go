package setup

import (
	"fmt"
	"slices"
	"strings"

	"golang.org/x/sys/windows/registry"

	charmlog "github.com/charmbracelet/log"
)

func appendPath(location string, path string) string {
	if len(location) > 0 {
		return location + ";" + path
	}
	return path
}

func readOnlyKey() (registry.Key, error) {
	return registry.OpenKey(registry.CURRENT_USER, "Environment", registry.QUERY_VALUE)
}

func writeKey() (registry.Key, error) {
	return registry.OpenKey(registry.CURRENT_USER, "Environment", registry.SET_VALUE)
}

func RemoveEnvVar(location string) error {
	// Remove the provided env location
	if location != "" {
		kr, err := readOnlyKey()
		if err != nil {
			charmlog.Error("Failed to open registry key to read", "error", err)
			return err
		}
		defer kr.Close()
		path, _, err := kr.GetStringValue("Path")
		if err != nil {
			charmlog.Error("Failed to get paths", "error", err)
			return err
		}
		paths := strings.Split(path, ";")

		for i, p := range paths {
			if p == location {
				paths = slices.Delete(paths, i, i+1)
			}
		}
		newPath := strings.Join(paths, ";")

		kw, err := writeKey()
		if err != nil {
			charmlog.Error("Failed to open registry key to write", "error", err)
			return err
		}
		defer kw.Close()
		err = kw.SetExpandStringValue("Path", newPath)
		if err != nil {
			charmlog.Error("Failed to set new user path", "error", err)
			return err
		}
		charmlog.Info("Path deleted from CURRENT_USER environment variables", "location", location)
	} else {
		return fmt.Errorf("Failed to remove environment variables (location empty)")
	}
	return nil
}

func RegisterEnvVar(location string) {
	kr, err := readOnlyKey()
	if err != nil {
		charmlog.Error("Failed to open registry key to read", "error", err)
		return
	}
	defer kr.Close()
	path, _, err := kr.GetStringValue("Path")
	if err != nil {
		charmlog.Error("Failed to get path", "error", err)
		return
	}
	newPath := appendPath(location, path)
	kw, err := writeKey()
	if err != nil {
		charmlog.Error("Failed to open registry key to write", "error", err)
		return
	}
	defer kw.Close()
	err = kw.SetExpandStringValue("Path", newPath)
	if err != nil {
		charmlog.Error("Failed to set new user path", "error", err)
		return
	}
	charmlog.Info("Path registerd into CURRENT_USER environment variables", "location", location)
}
