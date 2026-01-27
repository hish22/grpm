package configc

import (
	"fmt"
	"hish22/grpm/internal/config"

	"github.com/spf13/cobra"
)

var (
	show bool
)

func Config() *cobra.Command {
	c := &cobra.Command{
		Use:   "config",
		Short: "grpm configuration information",
		Run:   ConfigCmd,
	}
	c.Flags().BoolVarP(&show, "show", "s", false, "Show TOML configuration information")

	return c
}

func ConfigCmd(cmd *cobra.Command, args []string) {
	if show {
		c := config.DecodeTOMLConfig()
		fmt.Println("==========grpm Configuration==========")
		fmt.Println("Location:", "\033[1m", c.Location, "\033[0m", "=> Where your installed files are saved")
	}
}
