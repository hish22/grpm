package commands

import (
	"hish22/grpm/cmd/execution"
	"hish22/grpm/cmd/flags"

	charmlog "github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

// func printConfigInformation() error {
// 	c, err := config.DecodeTOMLConfig()
// 	if err != nil {
// 		return err
// 	}
// 	fmt.Println("==========grpm Configuration==========")
// 	fmt.Println("Location:", "\033[1m", c.Location, "\033[0m", "=> Main grpm directory")
// 	fmt.Println("Downloaded:", "\033[1m", c.Downloaded, "\033[0m", "=> Where your downloaded files are saved")
// 	fmt.Println("library location:", "\033[1m", c.Library, "\033[0m", "=> Where your installed files are saved (files downloaded with --setup)")
// 	fmt.Println("System Architecture:", "\033[1m", c.Arch, "\033[0m", "=> Your system architecture (ex: x64,amd64,aarch64,etc)")
// 	fmt.Println("Operating System:", "\033[1m", c.Os, "\033[0m", "=> Your own operating system")
// 	return nil
// }

// func openConfigFile() error {
// 	return config.OpenTOMLConfig()
// }

func ConfigC() *cobra.Command {
	c := &cobra.Command{
		Use:   "config",
		Short: "grpm configuration information",
		RunE: func(cmd *cobra.Command, args []string) error {
			e, err := execution.NewExecution(
				"configc",
				[]execution.AtomicExecution{
					// execution.UnsafeAtomicExecution("print_config", printConfigInformation, *flags.GetBooleanFlag("show")),
					// execution.UnsafeAtomicExecution("open_config", openConfigFile, *flags.GetBooleanFlag("rfconfig")),
				},
				func() error {
					err := cmd.Help()
					if err != nil {
						charmlog.Error("Failed to display grpm configuration help list,", "error", err)
					}
					return err
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
	flags.SetBooleanFlag("show", c.Flags().BoolP("show", "s", false, "retrive grpm configuration information."))
	flags.SetBooleanFlag("rfconfig", c.Flags().BoolP("open", "o", false, "open configuration file for edit."))
	return c
}
