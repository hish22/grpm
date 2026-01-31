package searchc

import (
	"fmt"
	"hish22/grpm/internal/packet"
	"hish22/grpm/internal/search"
	"log"
	"strconv"

	"github.com/spf13/cobra"
)

var (
	repo      string
	page      int
	mostStars bool
	fewStars  bool
)

func Search() *cobra.Command {
	c := &cobra.Command{
		Use:   "search",
		Short: "Search a specific github object.",
		Run:   searchCmd,
	}
	c.Flags().StringVarP(&repo, "repo", "r", "", "Search for specific repo")
	c.Flags().IntVarP(&page, "page", "p", 1, "Specifiy page for a repo")
	c.Flags().BoolVarP(&mostStars, "most-stars", "m", false, "Search for most stars repo")
	c.Flags().BoolVarP(&fewStars, "few-stars", "f", false, "Search for fewer stars repo")
	return c
}

func searchCmd(cmd *cobra.Command, args []string) {
	if len(repo) != 0 {
		repoSearch()
	} else {
		if err := cmd.Help(); err != nil {
			log.Fatal("Can't print search help command,", err)
		}
	}
}

func repoSearch() {
	HitRepos, err := search.JsonSearchRepo(&packet.RepoInfo{
		Name:      repo,
		Page:      strconv.Itoa(page),
		MostStars: mostStars,
		FewStars:  fewStars,
	})

	if err != nil {
		log.Fatal("(grpm search) didn't find any repository")
	}

	if len(HitRepos) != 0 {
		for _, r := range HitRepos {
			fmt.Printf("\n\033]8;;https://github.com/%s/%s\a\033[1m%s/%s (%s stars)\033[0m\033]8;;\a\n%s\n",
				r.Owner, r.Name, r.Owner, r.Name, r.Stars, r.Description)
			fmt.Println() // last space
		}
	} else {
		fmt.Println("\033[1mNo result found of", repo, "\033[0m")
	}
}
