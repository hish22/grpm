package listc

import (
	"hish22/grpm/internal/list"

	"github.com/spf13/cobra"
)

func ListC() *cobra.Command {
	c := &cobra.Command{
		Use:   "list",
		Short: "List installed assets",
		Run:   listCmd,
	}
	return c
}

func listCmd(cmd *cobra.Command, args []string) {
	list.ListAssets()
}
