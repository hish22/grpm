package configc

import (
	"fmt"
	"hish22/grpm/internal/config"

	"github.com/spf13/cobra"
)

var (
	show     bool // Print TOML config information
	rfconfig bool // Allow refactor
	//rconfig   string // Refactor config type (like location)
	//crname    string // Refactor value of the config type (like location = "/usr/bin")
)

func Config() *cobra.Command {
	c := &cobra.Command{
		Use:   "config",
		Short: "grpm configuration information",
		Run:   ConfigCmd,
	}
	c.Flags().BoolVarP(&show, "show", "s", false, "Show TOML configuration information")
	c.Flags().BoolVarP(&rfconfig, "open", "o", false, "Open TOML configuration file")
	return c
}

func ConfigCmd(cmd *cobra.Command, args []string) {
	c := config.DecodeTOMLConfig()
	if show {
		fmt.Println("==========grpm Configuration==========")
		fmt.Println("Location:", "\033[1m", c.Location, "\033[0m", "=> Where your installed files are saved")
	} else if rfconfig {
		config.OpenTOMLConfig()
	}
}
