package infoc

import (
	"fmt"
	"hish22/grpm/internal/info"
	"hish22/grpm/internal/packet"
	"hish22/grpm/internal/search"
	"log"

	"github.com/spf13/cobra"
)

var (
	rName     string
	ownerName string
)

func Info() *cobra.Command {
	c := &cobra.Command{
		Use:   "info",
		Short: "display repository's information",
		Run:   infoCmd,
	}
	c.Flags().StringVarP(&ownerName, "owner", "o", "", "Repository owner")
	c.Flags().StringVarP(&rName, "repo", "r", "", "Repository name")
	return c
}

func infoCmd(cmd *cobra.Command, args []string) {
	var pInfo packet.RepoPageInfo
	var err error
	if len(rName) != 0 && len(ownerName) != 0 {
		pInfo, err = info.JsonInfoRepo(&ownerName, &rName)

		if err != nil {
			log.Fatal("(grpm info) didn't find any repository information")
		}

	} else {
		// In case a user didn't specifiy a repo name or owner
		// we will perfom a search command and use first
		// most stars result as the entry to search for
		repo := packet.RepoInfo{}
		repo.FewStars = false
		repo.MostStars = true
		repo.Page = "1"
		if len(rName) == 0 && len(ownerName) != 0 {
			repo.Name = ownerName
		} else if len(ownerName) == 0 && len(rName) != 0 {
			repo.Name = rName
		} else {
			cmd.Help()
			return
		}

		SearchRepo, error := search.JsonSearchRepo(&repo)

		if error != nil {
			log.Fatal("(grpm search) didn't find any repository")
		}

		pInfo, err = info.JsonInfoRepo(&SearchRepo[0].Owner, &SearchRepo[0].Name)

		if err != nil {
			log.Fatal("(grpm info) didn't find any repository information")
		}
	}
	fmt.Println("=== Repository Information ===")
	fmt.Println("ID: ", pInfo.ID)
	fmt.Println("Owner: ", pInfo.Owner)
	fmt.Println("Repository Name: ", pInfo.RepoName)
	fmt.Println("Created at: ", pInfo.CreatedAt)
}
