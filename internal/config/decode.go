package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

func DecodeTOMLConfig() config {

	var conf config

	_, err := toml.Decode(fetchTOMLconfig(), &conf)
	if err != nil {
		log.Fatal("Can't decode TOML config file, ", err)
	}

	return conf
}
