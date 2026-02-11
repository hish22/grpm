package updatec

import (
	"hish22/grpm/internal/update"

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
