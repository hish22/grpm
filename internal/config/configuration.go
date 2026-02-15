package config

type config struct {
	Downloaded string `toml:"downloaded"`
	Library    string `toml:"library"`
	Location   string `toml:"location"`
	Arch       string `toml:"arch"`
	Os         string `toml:"os"`
}
