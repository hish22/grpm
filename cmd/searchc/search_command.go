package searchc

import (
	"fmt"
	"hish22/grpm/cmd/flags"
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
<<<<<<< HEAD
		Short: "Search github repositories.",
=======
		Short: "Search for set of repositories",
>>>>>>> master
		Run:   searchCmd,
		PreRun: func(cmd *cobra.Command, args []string) {
			if !config.CheckConfig() {
				charmlog.Fatal("Please Run (grpm -d) to define grpm configuration files")
			}
		},
	}
<<<<<<< HEAD
	flags.StringFlags["repo_name"] = c.Flags().StringP("name", "r", "", "repository name to search for.")
	flags.IntgerFlags["repo_search_page"] = c.Flags().IntP("page", "p", 1, "page number (Default 1).")
	flags.StringFlags["repo_search_sort"] = c.Flags().StringP("sort", "s", "", "sort repositories based on provided criteria (stars, forks, help-wanted-issues, updated).")
	flags.StringFlags["repo_search_order"] = c.Flags().StringP("order", "o", "", "Specify direction: 'asc' (A-Z/oldest first) or 'desc' (Z-A/newest first).")
	// c.Flags().StringVarP(&repo, "repo", "r", "", "Search a list of repositories.")
	// c.Flags().IntVarP(&page, "page", "p", 1, "page number of the results to fetch (Default 1).")
	// c.Flags().StringVarP(&sort, "sort", "s", "", "Sort repositories based criteria (stars, forks, help-wanted-issues, updated).")
	// c.Flags().StringVarP(&order, "order", "o", "", "Order of sorting repositories (asc, desc).")
=======
	c.Flags().StringVarP(&repo, "repo", "r", "", "Search a list of repositories.")
	c.Flags().IntVarP(&page, "page", "p", 1, "page number of the results to fetch (Default 1).")
	c.Flags().StringVarP(&sort, "sort", "s", "", "Sort repositories based on a criteria (stars, forks, help-wanted-issues, updated).")
	c.Flags().StringVarP(&order, "order", "o", "", "Order of sorting repositories (asc, desc).")
>>>>>>> master
	return c
}

func searchCmd(cmd *cobra.Command, args []string) {
	if len(repo) != 0 {
		s := &search.RepoInfo{
			Name:  *flags.StringFlags["repo_name"],
			Page:  strconv.Itoa(*flags.IntgerFlags["repo_search_page"]),
			Sort:  *flags.StringFlags["repo_search_sort"],
			Order: *flags.StringFlags["repo_search_order"],
		}
		repositories, err := search.SearchRepositories(s)
		if err != nil {
			charmlog.Error("Failed to search specified repositories", "error", err)
			return
		}
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
