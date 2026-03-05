package releasec

import (
	"fmt"
	"hish22/grpm/internal/config"
	"hish22/grpm/internal/release"

	charmlog "github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

var (
	repo          string
	latest        bool
	tag           string
	latestRelease bool
)

func ReleaseC() *cobra.Command {
	c := cobra.Command{
		Use:   "release",
		Short: "Info about repo's releases information",
		Run:   releaseCmd,
		PreRun: func(cmd *cobra.Command, args []string) {
			if !config.CheckConfig() {
				charmlog.Fatal("Please Run (grpm -d) to define grpm configuration files")
			}
		},
	}
	c.Flags().StringVarP(&repo, "repo", "r", "", "Repository name (owner/repo)")
	c.Flags().BoolVarP(&latest, "latest", "a", false, "Show 5 latest repository releases information")
	c.Flags().StringVarP(&tag, "tag", "t", "", "Show a specific repository release by a tag")
	c.Flags().BoolVarP(&latestRelease, "latest-release", "l", false, "Show latest repository release information")
	return &c
}

func releaseCmd(cmd *cobra.Command, args []string) {
	if latest && len(tag) == 0 {
		releases, err := release.FetchLatestReleases(&repo)
		if err != nil {
			charmlog.Error("Failed to fetch latest repository releases", "error", err)
			return
		}
		fmt.Println("=== Repo's Latest Releases ===")
		for _, r := range releases {
			fmt.Print("\n")
			fmt.Println("ID: ", r.ID)
			fmt.Println("Release Name: ", r.ReleaseName)
			fmt.Println("Tag Name: ", r.TagName)
			fmt.Println("Created At: ", r.CreatedAt)
			fmt.Println("Updated At: ", r.UpdatedAt)
			fmt.Println("Release Page: ", r.HtmlUrl)
			fmt.Print("\n")
		}
		return
	} else if len(tag) != 0 {
		release, err := release.FetchSpecificRelease(repo, tag)
		if err != nil {
			charmlog.Error("Failed to fetch specified repository release", "error", err)
			return
		}
		fmt.Println("=== Repo's Release ===")
		fmt.Print("\n")
		fmt.Println("ID: ", release.ID)
		fmt.Println("Release Name: ", release.ReleaseName)
		fmt.Println("Tag Name: ", release.TagName)
		fmt.Println("Created At: ", release.CreatedAt)
		fmt.Println("Updated At: ", release.UpdatedAt)
		fmt.Println("Release Page: ", release.HtmlUrl)
		fmt.Print("\n")
	} else if latestRelease {
		release, err := release.FetchLatestRelease(repo)
		if err != nil {
			charmlog.Error("Failed to fetch Latest repository release", "error", err)
			return
		}
		fmt.Println("=== Repo's Release ===")
		fmt.Print("\n")
		fmt.Println("ID: ", release.ID)
		fmt.Println("Release Name: ", release.ReleaseName)
		fmt.Println("Tag Name: ", release.TagName)
		fmt.Println("Created At: ", release.CreatedAt)
		fmt.Println("Updated At: ", release.UpdatedAt)
		fmt.Println("Release Page: ", release.HtmlUrl)
		fmt.Print("\n")
	} else {
		cmd.Help()
	}
}
