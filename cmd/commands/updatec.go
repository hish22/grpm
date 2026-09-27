package commands

import (
	"hish22/grpm/cmd/execution"
	"hish22/grpm/cmd/flags"

	"github.com/spf13/cobra"
)

// func updateCmd(cmd *cobra.Command, args []string) {
// 	if latest && len(repo) != 0 {
// 		update.UpdateToLatestAsset(repo, force)
// 	} else if check {
// 		update.CheckUpdate(repo)
// 	} else {
// 		cmd.Help()
// 	}
// }

// var (
// 	repo   string
// 	latest bool
// 	// setup  bool
// 	force bool
// 	check bool
// )

func UpdateC() *cobra.Command {
	c := &cobra.Command{
		Use:   "update",
		Short: "Update installed assets",
		RunE: func(cmd *cobra.Command, args []string) error {
			e, err := execution.NewExecution(
				"updatec",
				[]execution.AtomicExecution{},
				func() error {
					return nil
				},
			)

			if err != nil {
				return err
			}

			return e.SequentialRun(true)

		},
		// PreRun: func(cmd *cobra.Command, args []string) {
		// 	if !config.CheckConfig() {
		// 		charmlog.Fatal("Please Run (grpm -d) to define grpm configuration files")
		// 	}
		// },
	}
	flags.SetStringFlag("repo_name", c.Flags().StringP("repo", "r", "", "Repository name (Owner/repo)"))
	flags.SetBooleanFlag("latest", c.Flags().BoolP("latest", "l", false, "Update to latest asset"))
	flags.SetBooleanFlag("force", c.Flags().BoolP("force", "f", false, "Apply any confirmation message as yes"))
	flags.SetBooleanFlag("check", c.Flags().BoolP("check", "c", false, "Check if the specified asset has an update"))
	// c.Flags().StringVarP(&repo, "repo", "r", "", "Repository name (Owner/repo)")
	// c.Flags().BoolVarP(&latest, "latest", "l", false, "Update to latest asset")
	// c.Flags().BoolVarP(&force, "force", "f", false, "Apply any confirmation message as yes")
	// c.Flags().BoolVarP(&check, "check", "c", false, "Check if the specified asset has an update")
	return c
}
