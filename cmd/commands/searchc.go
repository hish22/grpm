package commands

import (
	"hish22/grpm/cmd/execution"
	"hish22/grpm/cmd/flags"

	"github.com/spf13/cobra"
)

// func searchCmd(cmd *cobra.Command, args []string) {
// 	if len(*flags.StringFlags["repo_name"]) != 0 {
// 		s := &search.RepoInfo{
// 			Name:  *flags.StringFlags["repo_name"],
// 			Page:  strconv.Itoa(*flags.IntgerFlags["repo_search_page"]),
// 			Sort:  *flags.StringFlags["repo_search_sort"],
// 			Order: *flags.StringFlags["repo_search_order"],
// 		}
// 		repositories, err := search.SearchRepositories(s)
// 		if err != nil {
// 			charmlog.Error("Failed to search specified repositories", "error", err)
// 			return
// 		}
// 		enumerateRepos(repositories)
// 	} else {
// 		if err := cmd.Help(); err != nil {
// 			charmlog.Fatal("Failed to show search help opitions,", "Error", err)
// 		}
// 	}
// }

// func enumerateRepos(repositories *structures.Repositories) {
// 	if len(repositories.Repositories) > 0 {
// 		for _, r := range repositories.Repositories {
// 			fmt.Printf("\n\033]8;;https://github.com/%s/%s\a\033[1m%s/%s (%d stars | %d forks)\033[0m\033]8;;\a\n%s\n",
// 				r.Owner.Username, r.Name, r.Owner.Username, r.Name, r.Stars, r.Forks, r.Description)
// 			fmt.Println()
// 		}
// 	} else {
// 		fmt.Println("\033[1mNo result found of", repo, "\033[0m")
// 	}
// }

// var (
// 	repo  string
// 	order string
// 	sort  string
// 	page  int
// )

func SearchC() *cobra.Command {
	c := &cobra.Command{
		Use:   "search",
		Short: "Search github repositories.",
		RunE: func(cmd *cobra.Command, args []string) error {
			e, err := execution.NewExecution(
				"searchc",
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
	flags.SetStringFlag("name", c.Flags().StringP("name", "r", "", "repository name to search for."))
	flags.SetIntgerFlag("page", c.Flags().IntP("page", "p", 1, "page number (Default 1)."))
	flags.SetStringFlag("sort", c.Flags().StringP("sort", "s", "", "sort repositories based on provided criteria (stars, forks, help-wanted-issues, updated)."))
	flags.SetStringFlag("order", c.Flags().StringP("order", "o", "", "Specify direction: 'asc' (A-Z/oldest first) or 'desc' (Z-A/newest first)."))
	// c.Flags().StringVarP(&repo, "repo", "r", "", "Search a list of repositories.")
	// c.Flags().IntVarP(&page, "page", "p", 1, "page number of the results to fetch (Default 1).")
	// c.Flags().StringVarP(&sort, "sort", "s", "", "Sort repositories based criteria (stars, forks, help-wanted-issues, updated).")
	// c.Flags().StringVarP(&order, "order", "o", "", "Order of sorting repositories (asc, desc).")
	return c
}
