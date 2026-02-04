package infoc

import (
	"fmt"
	"hish22/grpm/internal/info"
	"hish22/grpm/internal/search"
	"hish22/grpm/internal/structures"

	"github.com/charmbracelet/lipgloss"
	charmlog "github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

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

var (
	name  string
	owner string
)

func InfoC() *cobra.Command {
	c := &cobra.Command{
		Use:   "info",
		Short: "display repository's information",
		Run:   infoCmd,
	}
	c.Flags().StringVarP(&owner, "owner", "o", "", "Repository owner")
	c.Flags().StringVarP(&name, "repo", "r", "", "Repository name")
	return c
}

func fullRepositoryInfo() {
	repository := info.InfoRepository(&owner, &name)
	showRepositoryInfo(repository)
}

func particalRepositoryInfo() {
	// In case a user didn't specifiy a repo name or owner
	// we will perfom a search command and use first
	// most stars result as the entry to search for
	repo := search.RepoInfo{}
	repo.Order = "desc"
	repo.Sort = "stars"
	repo.Page = "1"
	if len(name) == 0 && len(owner) != 0 {
		repo.Name = owner
	} else if len(owner) == 0 && len(name) != 0 {
		repo.Name = name
	}
	SearchedRepo := search.SearchRepositories(&repo)
	repository := info.InfoRepository(&SearchedRepo.Repositories[0].Owner.Username, &SearchedRepo.Repositories[0].Name)
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
	if len(name) != 0 && len(owner) != 0 {
		fullRepositoryInfo()
	} else if len(name) == 0 && len(owner) == 0 {
		err := cmd.Help()
		if err != nil {
			charmlog.Fatal("Failed to show info help", err)
		}
	} else {
		particalRepositoryInfo()
	}
}
