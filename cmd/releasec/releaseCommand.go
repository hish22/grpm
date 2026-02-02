package releasec

import (
	"fmt"
	"hish22/grpm/internal/release"

	"github.com/spf13/cobra"
)

var (
	repo   string
	latest bool
	tag    string
)

func ReleaseC() *cobra.Command {
	c := cobra.Command{
		Use:   "release",
		Short: "Info about repo's releases information",
		Run:   releaseCmd,
	}
	c.Flags().StringVarP(&repo, "repo", "r", "", "Repo's name (owner/repo)")
	c.Flags().BoolVarP(&latest, "latest", "l", false, "Grab 5 latest repo's releases information")
	c.Flags().StringVarP(&tag, "tag", "t", "", "Grab a specific repo's release by tag")
	return &c
}

func releaseCmd(cmd *cobra.Command, args []string) {
	if latest && len(tag) == 0 {
		releases := release.FetchLatestReleases(&repo)
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
		release := release.FetchSpecificRelease(&repo, &tag)
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
