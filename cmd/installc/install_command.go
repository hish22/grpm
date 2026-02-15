package installc

import (
	"bufio"
	"fmt"
	"hish22/grpm/internal/asset"
	"hish22/grpm/internal/config"
	"hish22/grpm/internal/install"
	"hish22/grpm/internal/release"
	"log"
	"os"
	"strconv"

	charmlog "github.com/charmbracelet/log"
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
		PreRun: func(cmd *cobra.Command, args []string) {
			if !config.CheckConfig() {
				charmlog.Fatal("Please Run (grpm -d) to define grpm configuration files")
			}
		},
	}
	c.Flags().StringVarP(&repo, "repo", "r", "", "Repo's name (owner/repo)")
	c.Flags().StringVarP(&tag, "tag", "t", "", "Grab a specific repo's release by tag")
	c.Flags().BoolVarP(&match, "match", "m", false, "Print assets that match your config file opitions")
	c.Flags().BoolVarP(&setup, "setup", "s", false, "Auto setup of installed asset")
	return &c
}

func scanner() int {
	s := bufio.NewScanner(os.Stdin)
	fmt.Print("Specifiy asset (index): ")
	if s.Scan() {
		index, err := strconv.Atoi(s.Text())
		if err != nil {
			charmlog.Fatal("No such selection (Not applicable)", "error", err)
		}
		return index
	}
	if s.Err() != nil {
		log.Fatal("Failed to scan the data, ", s.Err())
	}
	return 0
}

func installCmd(cmd *cobra.Command, args []string) {
	if len(repo) != 0 && len(tag) != 0 {
		a, err := release.FetchSpecificRelease(repo, tag)
		if err != nil {
			charmlog.Error("Failed to fetch specified repository release", "error", err)
			return
		}
		asset.PrintTheAssets(a, repo, match)
		ch := scanner()

		if ch > (len(a.Assets) - 1) {
			charmlog.Error("No such asset")
			return
		}

		chRelease := a.Assets[ch]
		install.InstallSelectedAsset(repo, &chRelease, a, setup)
	} else if len(tag) == 0 && len(repo) != 0 {
		a, err := release.FetchLatestRelease(repo)
		if err != nil {
			charmlog.Error("Failed to fetch latest repository release", "error", err)
			return
		}
		asset.PrintTheAssets(a, repo, match)
		ch := scanner()

		if ch > (len(a.Assets) - 1) {
			charmlog.Error("No such asset")
			return
		}

		chRelease := a.Assets[ch]
		install.InstallSelectedAsset(repo, &chRelease, a, setup)
	} else {
		cmd.Help()
	}
}
