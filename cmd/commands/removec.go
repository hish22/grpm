package commands

import (
	"hish22/grpm/cmd/execution"
	"hish22/grpm/cmd/flags"

	"github.com/spf13/cobra"
)

// func removeCmd(cmd *cobra.Command, args []string) {
// 	if repo != "" {
// 		remove.RemoveAssetByRepoName(repo)
// 	}
// }

// var (
// 	repo string
// )

func RemoveC() *cobra.Command {
	c := &cobra.Command{
		Use:   "remove",
		Short: "Remove an installed asset",
		RunE: func(cmd *cobra.Command, args []string) error {
			e, err := execution.NewExecution(
				"removec",
				[]execution.AtomicExecution{},
				func() error {
					return nil
				},
			)

			if err != nil {
				return nil
			}

			return e.SequentialRun(true)

		},
		// PreRun: func(cmd *cobra.Command, args []string) {
		// 	if !config.CheckConfig() {
		// 		charmlog.Fatal("Please Run (grpm -d) to define grpm configuration files")
		// 	}
		// },
	}
	flags.SetStringFlag("repo_name", c.Flags().StringP("repo", "r", "", "Repository name (owner/repo)"))
	// c.Flags().StringVarP(&repo, "repo", "r", "", "Repository name (owner/repo)")
	return c
}
