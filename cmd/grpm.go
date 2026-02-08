package main

import (
	"hish22/grpm/cmd/cachec"
	"hish22/grpm/cmd/configc"
	"hish22/grpm/cmd/infoc"
	"hish22/grpm/cmd/installc"
	"hish22/grpm/cmd/listc"
	"hish22/grpm/cmd/releasec"
	"hish22/grpm/cmd/searchc"
	"hish22/grpm/cmd/updatec"
	"hish22/grpm/internal/config"

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
		Run: func(cmd *cobra.Command, args []string) {
			if define {
				config.GenerateTOMLConfig()
			} else {
				cmd.Help()
			}
		},
	}
}

func main() {
	r := root()
	r.Flags().BoolVarP(&define, "define", "d", false, "initialize your grpm tool")
	// Add search command
	r.AddCommand(searchc.SearchC())
	r.AddCommand(configc.ConfigC())
	r.AddCommand(infoc.InfoC())
	r.AddCommand(releasec.ReleaseC())
	r.AddCommand(installc.InstallC())
	r.AddCommand(listc.ListC())
	r.AddCommand(updatec.UpdateC())
	r.AddCommand(cachec.CacheC())
	if err := r.Execute(); err != nil {
		panic(err)
	}
}
