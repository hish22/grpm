package cachec

import (
	"hish22/grpm/internal/persistance"

	"github.com/spf13/cobra"

	charmlog "github.com/charmbracelet/log"
)

var (
	clear bool
)

func CacheC() *cobra.Command {
	c := &cobra.Command{
		Use:   "cache",
		Short: "Handle cache commands",
		Run:   cacheCmd,
	}
	c.Flags().BoolVarP(&clear, "clear", "c", false, "Clear all stored cache")
	return c
}

func cacheCmd(cmd *cobra.Command, args []string) {
	if clear {
		persistance.ClearCache()
		charmlog.Info("Cache cleared")
	} else {
		err := cmd.Help()
		if err != nil {
			charmlog.Fatal("Failed to show cache help", "error", err)
		}
	}

}
