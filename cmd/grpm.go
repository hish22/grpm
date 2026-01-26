package main

import (
	"hish22/grpm/cmd/search"

	"github.com/spf13/cobra"
)

func root() *cobra.Command {
	return &cobra.Command{
		Use:   "grpm",
		Short: "A cool github release packet manger",
		Long:  `Github Releases Packet Manager (grpm) is a tool to handle installed releases from github.`,
	}
}

func main() {

	r := root()

	// Add search command
	r.AddCommand(search.Search())

	if err := r.Execute(); err != nil {
		panic(err)
	}
}
