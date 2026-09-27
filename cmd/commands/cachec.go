package commands

import (
	"hish22/grpm/cmd/execution"
	"hish22/grpm/cmd/flags"

	"github.com/spf13/cobra"

	charmlog "github.com/charmbracelet/log"
)

// func clearCache() error {
// 	if persistance.ClearCache() {
// 		charmlog.Info("Cache cleared.")
// 		return nil
// 	}
// 	return errors.New("Failed to clear cache.")
// }

func CacheC() *cobra.Command {
	c := &cobra.Command{
		Use:   "cache",
		Short: "Cache management command.",
		RunE: func(cmd *cobra.Command, args []string) error {
			e, err := execution.NewExecution(
				"cachec",
				[]execution.AtomicExecution{
					// execution.UnsafeAtomicExecution("clear_cache", clearCache, *flags.GetBooleanFlag("clear")),
				},
				func() error {
					err := cmd.Help()
					if err != nil {
						charmlog.Error("Failed to display cache management help list.", "error", err)
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
		// 		charmlog.Fatal("Please Run (grpm -d) to define grpm configuration files.")
		// 	}
		// },
	}
	flags.SetBooleanFlag("clear",
		c.Flags().BoolP("clear", "c", false, "clear searched and requested info results from cache."))
	return c
}
