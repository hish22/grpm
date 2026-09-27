package main

import (
	"hish22/grpm/cmd/commands"

	charmlog "github.com/charmbracelet/log"

	"github.com/spf13/cobra"
)

var (
	define bool
)

func root() *cobra.Command {
	return &cobra.Command{
		Use:   "grpm",
		Short: "A cool github release packet manger",
		Long:  `Github Releases Packet Manager (grpm) is a tool to handle installed releases from github.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
}

func main() {
	r := root()
	r.Flags().BoolVarP(&define, "define", "d", false, "Initialize grpm tool.")
	grpmCommands := []*cobra.Command{
		commands.SearchC(),
		commands.ConfigC(),
		commands.InfoC(),
		commands.ReleaseC(),
		commands.InstallC(),
		commands.ListC(),
		commands.UpdateC(),
		commands.CacheC(),
		commands.RemoveC(),
	}
	// Add search command
	r.AddCommand(grpmCommands...)
	if err := r.Execute(); err != nil {
		charmlog.Error(err)
	}
}
