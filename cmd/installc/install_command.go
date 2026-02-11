package installc

import (
	"bufio"
	"fmt"
	"hish22/grpm/internal/asset"
	"hish22/grpm/internal/install"
	"hish22/grpm/internal/release"
	"log"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

var (
	repo  string
	tag   string
	match bool
	setup bool
)

func InstallC() *cobra.Command {
	c := cobra.Command{
		Use:   "install",
		Short: "Install a release",
		Run:   installCmd,
	}
	c.Flags().StringVarP(&repo, "repo", "r", "", "Repo's name (owner/repo)")
	c.Flags().StringVarP(&tag, "tag", "t", "", "Grab a specific repo's release by tag")
	c.Flags().BoolVarP(&match, "match", "m", false, "Print assets that match your config file opitions")
	c.Flags().BoolVarP(&setup, "setup", "s", false, "Auto setup of installed asset")
	return &c
}

func scanner() int {
	s := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter your choose: ")
	if s.Scan() {
		index, err := strconv.Atoi(s.Text())
		if err != nil {
			log.Fatal("(Wrong entry) Can't convert from string to int, ", err)
		}
		return index
	}
	if s.Err() != nil {
		log.Fatal("Can't scan the data, ", s.Err())
	}
	return 0
}

func installCmd(cmd *cobra.Command, args []string) {
	if len(repo) != 0 && len(tag) != 0 {
		a := release.FetchSpecificRelease(repo, tag)
		asset.PrintTheAssets(a, repo, match)
		ch := scanner()

		if ch > (len(a.Assets) - 1) {
			fmt.Println("No such asset")
			return
		}

		chRelease := a.Assets[ch]
		install.InstallSelectedAsset(repo, &chRelease, a, setup)
	} else if len(tag) == 0 && len(repo) != 0 {
		a := release.FetchLatestRelease(repo)
		asset.PrintTheAssets(a, repo, match)
		ch := scanner()

		if ch > (len(a.Assets) - 1) {
			fmt.Println("No such asset")
			return
		}

		chRelease := a.Assets[ch]
		install.InstallSelectedAsset(repo, &chRelease, a, setup)
	} else {
		cmd.Help()
	}
}
