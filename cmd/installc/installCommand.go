package installc

import (
	"bufio"
	"fmt"
	"hish22/grpm/internal/install"
	"log"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

var (
	repo string
	tag  string
)

func InstallC() *cobra.Command {
	c := cobra.Command{
		Use:   "install",
		Short: "Install a release",
		Run:   installCmd,
	}
	c.Flags().StringVarP(&repo, "repo", "r", "", "Repo's name (owner/repo)")
	c.Flags().StringVarP(&tag, "tag", "t", "", "Grab a specific repo's release by tag")
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
		a, r := install.FetchAssets(&repo, &tag)
		ch := scanner()

		if ch > (len(a) - 1) {
			fmt.Println("No such asset")
			return
		}

		chRelease := a[ch]
		install.InstallSelectedAsset(&chRelease, r)
	} else if len(tag) == 0 && len(repo) != 0 {
		a, r := install.FetchLatestReleaseAssets(&repo)
		ch := scanner()

		if ch > (len(a) - 1) {
			fmt.Println("No such asset")
			return
		}

		chRelease := a[ch]
		install.InstallSelectedAsset(&chRelease, r)
	} else {
		cmd.Help()
	}
}
