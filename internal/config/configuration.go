package config

type config struct {
	Downloaded string `toml:"downloaded"`
	Location   string `toml:"location"`
	Arch       string `toml:"arch"`
	Os         string `toml:"os"`
}
