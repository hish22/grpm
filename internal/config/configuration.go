package config

type config struct {
	Location string `toml:"location"`
	Arch     string `toml:"arch"`
	Os       string `toml:"os"`
}
