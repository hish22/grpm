package config

import (
	charmlog "github.com/charmbracelet/log"

	"github.com/BurntSushi/toml"
)

func DecodeTOMLConfig() (config, error) {
	var conf config
	_, err := toml.Decode(fetchTOMLconfig(), &conf)
	if err != nil {
		charmlog.Error("Failed to display configuration information", "error", err)
		// charmlog.Error("Unable to decode TOML config file", "error", err)
	}
	return conf, err
}
