package commands

import (
	"hish22/grpm/cmd/execution"
	"hish22/grpm/cmd/flags"

	"github.com/spf13/cobra"
)

// func releaseCmd(cmd *cobra.Command, args []string) {
// 	if latest && len(tag) == 0 {
// 		releases, err := release.FetchLatestReleases(&repo)
// 		if err != nil {
// 			charmlog.Error("Failed to fetch latest repository releases", "error", err)
// 			return
// 		}
// 		fmt.Println("=== Repo's Latest Releases ===")
// 		for _, r := range releases {
// 			fmt.Print("\n")
// 			fmt.Println("ID: ", r.ID)
// 			fmt.Println("Release Name: ", r.ReleaseName)
// 			fmt.Println("Tag Name: ", r.TagName)
// 			fmt.Println("Created At: ", r.CreatedAt)
// 			fmt.Println("Updated At: ", r.UpdatedAt)
// 			fmt.Println("Release Page: ", r.HtmlUrl)
// 			fmt.Print("\n")
// 		}
// 		return
// 	} else if len(tag) != 0 {
// 		release, err := release.FetchSpecificRelease(repo, tag)
// 		if err != nil {
// 			charmlog.Error("Failed to fetch specified repository release", "error", err)
// 			return
// 		}
// 		fmt.Println("=== Repo's Release ===")
// 		fmt.Print("\n")
// 		fmt.Println("ID: ", release.ID)
// 		fmt.Println("Release Name: ", release.ReleaseName)
// 		fmt.Println("Tag Name: ", release.TagName)
// 		fmt.Println("Created At: ", release.CreatedAt)
// 		fmt.Println("Updated At: ", release.UpdatedAt)
// 		fmt.Println("Release Page: ", release.HtmlUrl)
// 		fmt.Print("\n")
// 	} else if latestRelease {
// 		release, err := release.FetchLatestRelease(repo)
// 		if err != nil {
// 			charmlog.Error("Failed to fetch Latest repository release", "error", err)
// 			return
// 		}
// 		fmt.Println("=== Repo's Release ===")
// 		fmt.Print("\n")
// 		fmt.Println("ID: ", release.ID)
// 		fmt.Println("Release Name: ", release.ReleaseName)
// 		fmt.Println("Tag Name: ", release.TagName)
// 		fmt.Println("Created At: ", release.CreatedAt)
// 		fmt.Println("Updated At: ", release.UpdatedAt)
// 		fmt.Println("Release Page: ", release.HtmlUrl)
// 		fmt.Print("\n")
// 	} else {
// 		cmd.Help()
// 	}
// }

// var (
// 	repo          string
// 	latest        bool
// 	tag           string
// 	latestRelease bool
// )

func ReleaseC() *cobra.Command {
	c := cobra.Command{
		Use:   "release",
		Short: "Repository releases information",
		RunE: func(cmd *cobra.Command, args []string) error {
			e, err := execution.NewExecution(
				"releasec",
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
	flags.SetStringFlag("repo_name", c.Flags().StringP("repo", "r", "", "Repository name (owner/repo)"))
	flags.SetStringFlag("repo_tag", c.Flags().StringP("tag", "t", "", "Show a specific repository release by a tag"))
	flags.SetBooleanFlag("latest", c.Flags().BoolP("latest", "a", false, "Show 5 latest repository releases information"))
	flags.SetBooleanFlag("latest-release", c.Flags().BoolP("latest-release", "l", false, "Show latest repository release information"))
	// c.Flags().StringVarP(&repo, "repo", "r", "", "Repository name (owner/repo)")
	// c.Flags().BoolVarP(&latest, "latest", "a", false, "Show 5 latest repository releases information")
	// c.Flags().StringVarP(&tag, "tag", "t", "", "Show a specific repository release by a tag")
	// c.Flags().BoolVarP(&latestRelease, "latest-release", "l", false, "Show latest repository release information")
	return &c
}
