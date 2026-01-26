package search

import (
	"fmt"
	"hish22/grpm/internal/search"

	"github.com/spf13/cobra"
)

var repo bool

func Search() *cobra.Command {
	c := &cobra.Command{
		Use:   "search",
		Short: "Search a specific github object.",
		Run:   searchUtil,
	}
	c.Flags().BoolVarP(&repo, "repo", "r", false, "Search for specific repo")
	return c
}

func searchUtil(cmd *cobra.Command, args []string) {
	if repo {
		HitRepos := search.SearchRepo(args[0], args[0])
		for _, r := range HitRepos {
			fmt.Println("\n\033[1m", r.Name, "\033[0m\n", r.Description)
			fmt.Println()
		}
	}
	fmt.Println("Specifiy a flag to start searching.")
}
