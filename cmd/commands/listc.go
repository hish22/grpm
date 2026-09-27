package commands

import (
	"hish22/grpm/cmd/execution"

	"github.com/spf13/cobra"
)

// func listCmd() error {
// 	list.ListAssets()
// 	return nil
// }

func ListC() *cobra.Command {
	c := &cobra.Command{
		Use:   "list",
		Short: "List installed assets",
		RunE: func(cmd *cobra.Command, args []string) error {
			e, err := execution.NewExecution(
				"listc",
				[]execution.AtomicExecution{
					// execution.UnsafeAtomicExecution("list_assets", listCmd, true),
				},
				func() error {
					err := cmd.Help()
					if err != nil {
						return err
					}
					return nil
				},
			)

			if err != nil {
				return err
			}

			return e.SequentialRun(false)

		},
		// PreRun: func(cmd *cobra.Command, args []string) {
		// 	if !config.CheckConfig() {
		// 		charmlog.Fatal("Please Run (grpm -d) to define grpm configuration files")
		// 	}
		// },
	}
	return c
}
