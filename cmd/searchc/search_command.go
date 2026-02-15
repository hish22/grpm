package searchc

import (
	"fmt"
	"hish22/grpm/internal/config"
	"hish22/grpm/internal/search"
	"hish22/grpm/internal/structures"
	"strconv"

	charmlog "github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

var (
	repo  string
	order string
	sort  string
	page  int
)

func SearchC() *cobra.Command {
	c := &cobra.Command{
		Use:   "search",
		Short: "Search a specific github object.",
		Run:   searchCmd,
		PreRun: func(cmd *cobra.Command, args []string) {
			if !config.CheckConfig() {
				charmlog.Fatal("Please Run (grpm -d) to define grpm configuration files")
			}
		},
	}
	c.Flags().StringVarP(&repo, "repo", "r", "", "Search a list of repositories.")
	c.Flags().IntVarP(&page, "page", "p", 1, "page number of the results to fetch (Default 1).")
	c.Flags().StringVarP(&sort, "sort", "s", "", "Sort repositories based criteria (stars, forks, help-wanted-issues, updated).")
	c.Flags().StringVarP(&order, "order", "o", "", "Order of sorting repositories (asc, desc).")
	return c
}

func searchCmd(cmd *cobra.Command, args []string) {
	if len(repo) != 0 {
		s := &search.RepoInfo{
			Name:  repo,
			Page:  strconv.Itoa(page),
			Sort:  sort,
			Order: order,
		}
		repositories := search.SearchRepositories(s)
		enumerateRepos(repositories)
	} else {
		if err := cmd.Help(); err != nil {
			charmlog.Fatal("Failed to show search help opitions,", "Error", err)
		}
	}
}

func enumerateRepos(repositories *structures.Repositories) {
	if len(repositories.Repositories) > 0 {
		for _, r := range repositories.Repositories {
			fmt.Printf("\n\033]8;;https://github.com/%s/%s\a\033[1m%s/%s (%d stars | %d forks)\033[0m\033]8;;\a\n%s\n",
				r.Owner.Username, r.Name, r.Owner.Username, r.Name, r.Stars, r.Forks, r.Description)
			fmt.Println()
		}
	} else {
		fmt.Println("\033[1mNo result found of", repo, "\033[0m")
	}
}
