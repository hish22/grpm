package infoc

import (
	"fmt"
	"hish22/grpm/cmd/flags"
	"hish22/grpm/internal/config"
	"hish22/grpm/internal/info"
	"hish22/grpm/internal/search"
	"hish22/grpm/internal/structures"

	"github.com/charmbracelet/lipgloss"
	charmlog "github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

// TODO: Spereate into different package (style)
var (
	// Style for labels (the keys)
	labelStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("86")).Bold(true)
	// Style for the main headers
	headerStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("205")).
			Bold(true).
			Underline(true).
			MarginTop(1).
			MarginBottom(1)
	// A subtle style for the values
	valueStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("252"))
	// Bordered box for the owner info to nest it visually
	boxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("62")).
			Padding(0, 1)
)

// var (
// 	name  string
// 	owner string
// )

func InfoC() *cobra.Command {
	c := &cobra.Command{
		Use:   "info",
		Short: "display repository's information",
		Run:   infoCmd,
		PreRun: func(cmd *cobra.Command, args []string) {
			if !config.CheckConfig() {
				charmlog.Fatal("Please Run (grpm -d) to define grpm configuration files")
			}
		},
	}
	flags.SetStringFlag("repo_name", c.Flags().StringP("repo", "r", "", "repository name."))
	flags.SetStringFlag("repo_owner", c.Flags().StringP("owner", "o", "", "repositroy owner."))
	return c
}

func fullRepositoryInfo() {
	repo := *flags.GetStringFlag("repo_name")
	owner := *flags.GetStringFlag("repo_owner")
	repository, err := info.InfoRepository(repo, owner)
	if err != nil {
		charmlog.Error("Failed to info specified repository")
		return
	}
	showRepositoryInfo(repository)
}

func particalRepositoryInfo() {
	// In case a user didn't specifiy a repo name or owner
	// we will perfom a search command and use first
	// most stars result as the entry to search for
	repo_name := *flags.GetStringFlag("repo_name")
	repo_owner := *flags.GetStringFlag("repo_owner")
	repo := search.RepoInfo{}
	repo.Order = "desc"
	repo.Sort = "stars"
	repo.Page = "1"
	if len(repo_name) == 0 && len(repo_owner) != 0 {
		repo.Name = repo_owner
	} else if len(repo_owner) == 0 && len(repo_name) != 0 {
		repo.Name = repo_name
	}
	SearchedRepo, err := search.SearchRepositories(&repo)
	if err != nil {
		charmlog.Error("Failed to search specified repositories", "error", err)
		return
	}
	repository, err := info.InfoRepository(SearchedRepo.Repositories[0].Owner.Username, SearchedRepo.Repositories[0].Name)
	if err != nil {
		charmlog.Error("Failed to info specified repository", "error", err)
		return
	}
	showRepositoryInfo(repository)
}

func showRepositoryInfo(repository *structures.Repository) {
	fmt.Println(headerStyle.Render("📦 REPOSITORY DETAILS"))

	// Helper to print styled lines
	printStat := func(label, value any) {
		fmt.Printf("%s %s\n", labelStyle.Render(fmt.Sprintf("%-20s", label)), valueStyle.Render(fmt.Sprint(value)))
	}

	printStat("ID:", repository.ID)
	printStat("Name:", repository.Name)
	printStat("Description:", repository.Description)

	fmt.Println(lipgloss.NewStyle().Foreground(lipgloss.Color("240")).Render("--- Stats ---"))

	printStat("Forks:", repository.Forks)
	printStat("Stars:", repository.Stars)
	printStat("Language:", repository.ProgrammingLanguage)
	printStat("License:", repository.License.Name)
	printStat("Link:", repository.HtmlUrl)

	// Show owner inside a box
	fmt.Println(headerStyle.Render("👤 OWNER"))
	showOwnerInfo(&repository.Owner)
}

func showOwnerInfo(owner *structures.Owner) {
	// Build the string first so we can wrap it in a box
	content := fmt.Sprintf("%s %s\n%s %s\n%s %s",
		labelStyle.Render("ID:   "), valueStyle.Render(fmt.Sprint(owner.ID)),
		labelStyle.Render("User: "), valueStyle.Render(owner.Username),
		labelStyle.Render("Page: "), valueStyle.Render(owner.HtmlUrl),
	)

	fmt.Println(boxStyle.Render(content))
}

func infoCmd(cmd *cobra.Command, args []string) {

	repo_name := *flags.GetStringFlag("repo_name")
	repo_owner := *flags.GetStringFlag("repo_owner")

	if len(repo_name) != 0 && len(repo_owner) != 0 {
		fullRepositoryInfo()
	} else if len(repo_name) == 0 && len(repo_owner) == 0 {
		err := cmd.Help()
		if err != nil {
			charmlog.Fatal("Failed to show info help", err)
		}
	} else {
		particalRepositoryInfo()
	}
}
