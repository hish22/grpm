package commands

import (
	"hish22/grpm/cmd/execution"
	"hish22/grpm/cmd/flags"

	"github.com/spf13/cobra"
)

// func scanner() int {
// 	s := bufio.NewScanner(os.Stdin)
// 	fmt.Print("Specifiy asset (index): ")
// 	if s.Scan() {
// 		index, err := strconv.Atoi(s.Text())
// 		if err != nil {
// 			charmlog.Fatal("No such selection (Not applicable)", "error", err)
// 		}
// 		return index
// 	}
// 	if s.Err() != nil {
// 		log.Fatal("Failed to scan the data, ", s.Err())
// 	}
// 	return 0
// }

// func installLatestRelease() error {
// 	a, err := release.FetchLatestRelease(*flags.GetStringFlag("repo"))
// 	if err != nil {
// 		charmlog.Error("Failed to fetch latest repository release", "error", err)
// 		return err
// 	}
// 	asset.PrintTheAssets(a, *flags.GetStringFlag("repo"), *flags.GetBooleanFlag("match"))
// 	ch := scanner()

// 	if ch > (len(a.Assets) - 1) {
// 		charmlog.Error("No such asset")
// 		return errors.New("No such asset")
// 	}

// 	chRelease := a.Assets[ch]
// 	install.InstallSelectedAsset(*flags.GetStringFlag("repo"), &chRelease, a, *flags.GetBooleanFlag("setup"), *flags.GetBooleanFlag("force"))
// 	return nil
// }

// func installSpecificRelease() error {
// 	a, err := release.FetchSpecificRelease(*flags.GetStringFlag("repo"), *flags.GetStringFlag("tag"))
// 	if err != nil {
// 		charmlog.Error("Failed to fetch specified repository release", "error", err)
// 		return err
// 	}
// 	asset.PrintTheAssets(a, *flags.GetStringFlag("repo"), *flags.GetBooleanFlag("match"))
// 	ch := scanner()

// 	if ch > (len(a.Assets) - 1) {
// 		charmlog.Error("No such asset")
// 		return err
// 	}

// 	chRelease := a.Assets[ch]
// 	install.InstallSelectedAsset(*flags.GetStringFlag("repo"), &chRelease, a, *flags.GetBooleanFlag("setup"), *flags.GetBooleanFlag("force"))
// 	return nil
// }

// func installCmd(cmd *cobra.Command, args []string) {

// 	if len(repo) != 0 && len(tag) != 0 {
// 		a, err := release.FetchSpecificRelease(repo, tag)
// 		if err != nil {
// 			charmlog.Error("Failed to fetch specified repository release", "error", err)
// 			return
// 		}
// 		asset.PrintTheAssets(a, repo, match)
// 		ch := scanner()

// 		if ch > (len(a.Assets) - 1) {
// 			charmlog.Error("No such asset")
// 			return
// 		}

// 		chRelease := a.Assets[ch]
// 		install.InstallSelectedAsset(repo, &chRelease, a, setup, force)
// 	} else if len(tag) == 0 && len(repo) != 0 {
// 		a, err := release.FetchLatestRelease(repo)
// 		if err != nil {
// 			charmlog.Error("Failed to fetch latest repository release", "error", err)
// 			return
// 		}
// 		asset.PrintTheAssets(a, repo, match)
// 		ch := scanner()

// 		if ch > (len(a.Assets) - 1) {
// 			charmlog.Error("No such asset")
// 			return
// 		}

// 		chRelease := a.Assets[ch]
// 		install.InstallSelectedAsset(repo, &chRelease, a, setup, force)
// 	} else {
// 		cmd.Help()
// 	}
// }

// var (
// 	repo  string
// 	tag   string
// 	match bool
// 	setup bool
// 	force bool
// )

func InstallC() *cobra.Command {
	c := cobra.Command{
		Use:   "install",
		Short: "Install a release asset",
		RunE: func(cmd *cobra.Command, args []string) error {
			e, err := execution.NewExecution(
				"installc",
				[]execution.AtomicExecution{
					// execution.UnsafeAtomicExecution("install_latest", installLatestRelease, len(*flags.GetStringFlag("tag")) == 0 && len(*flags.GetStringFlag("repo")) != 0),
					// execution.UnsafeAtomicExecution("install_specific", installSpecificRelease, len(*flags.GetStringFlag("repo")) != 0 && len(*flags.GetStringFlag("tag")) != 0),
				},
				func() error {
					err := cmd.Help()
					if err != nil {
						return err
					}
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
	flags.SetStringFlag("repo_tag", c.Flags().StringP("tag", "t", "", "Grab a specific repository release by a tag"))
	flags.SetBooleanFlag("match", c.Flags().BoolP("match", "m", false, "Show assets that match your config file opitions"))
	flags.SetBooleanFlag("setup", c.Flags().BoolP("setup", "s", false, "Auto setup of downloaded asset"))
	flags.SetBooleanFlag("force", c.Flags().BoolP("force", "f", false, "Apply any confirmation message as yes"))
	// c.Flags().StringVarP(&repo, "repo", "r", "", "Repository name (owner/repo)")
	// c.Flags().StringVarP(&tag, "tag", "t", "", "Grab a specific repository release by a tag")
	// c.Flags().BoolVarP(&match, "match", "m", false, "Show assets that match your config file opitions")
	// c.Flags().BoolVarP(&setup, "setup", "s", false, "Auto setup of downloaded asset")
	// c.Flags().BoolVarP(&force, "force", "f", false, "Apply any confirmation message as yes")
	return &c
}
