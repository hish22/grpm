package listc

import (
	"hish22/grpm/internal/config"
	"hish22/grpm/internal/list"

	charmlog "github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

func ListC() *cobra.Command {
	c := &cobra.Command{
		Use:   "list",
		Short: "List installed assets",
		Run:   listCmd,
		PreRun: func(cmd *cobra.Command, args []string) {
			if !config.CheckConfig() {
				charmlog.Fatal("Please Run (grpm -d) to define grpm configuration files")
			}
		},
	}
	return c
}

func listCmd(cmd *cobra.Command, args []string) {
	list.ListAssets()
}
