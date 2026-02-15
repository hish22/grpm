package updatec

import (
	"hish22/grpm/internal/config"
	"hish22/grpm/internal/update"

	charmlog "github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

var (
	repo   string
	latest bool
	setup  bool
)

func UpdateC() *cobra.Command {
	c := &cobra.Command{
		Use:   "update",
		Short: "Update installed assets",
		Run:   updateCmd,
		PreRun: func(cmd *cobra.Command, args []string) {
			if !config.CheckConfig() {
				charmlog.Fatal("Please Run (grpm -d) to define grpm configuration files")
			}
		},
	}
	c.Flags().StringVarP(&repo, "repo", "r", "", "Repository name (Owner/repo)")
	c.Flags().BoolVarP(&latest, "latest", "l", false, "Update to latest asset")
	return c
}

func updateCmd(cmd *cobra.Command, args []string) {
	if latest && len(repo) != 0 {
		update.UpdateToLatestAsset(repo)
	} else {
		cmd.Help()
	}
}
