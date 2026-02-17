package removec

import (
	"hish22/grpm/internal/config"
	"hish22/grpm/internal/remove"

	charmlog "github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

var (
	repo string
)

func RemoveC() *cobra.Command {
	c := &cobra.Command{
		Use:   "remove",
		Short: "Delete an installed asset",
		Run:   removeCmd,
		PreRun: func(cmd *cobra.Command, args []string) {
			if !config.CheckConfig() {
				charmlog.Fatal("Please Run (grpm -d) to define grpm configuration files")
			}
		},
	}
	c.Flags().StringVarP(&repo, "repo", "r", "", "Repository name")
	return c
}

func removeCmd(cmd *cobra.Command, args []string) {
	if repo != "" {
		remove.RemoveAssetByRepoName(repo)
	}
}
