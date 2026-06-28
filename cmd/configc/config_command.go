package configc

import (
	"fmt"
	"hish22/grpm/internal/config"
	"log"

	charmlog "github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

var (
	show     bool // Print TOML config information
	rfconfig bool // Allow refactor
	//rconfig   string // Refactor config type (like location)
	//crname    string // Refactor value of the config type (like location = "/usr/bin")
)

func ConfigC() *cobra.Command {
	c := &cobra.Command{
		Use:   "config",
		Short: "grpm configuration information",
		Run:   ConfigCmd,
		PreRun: func(cmd *cobra.Command, args []string) {
			if !config.CheckConfig() {
				charmlog.Fatal("Please Run (grpm -d) to define grpm configuration files")
			}
		},
	}
	c.Flags().BoolVarP(&show, "show", "s", false, "Show TOML configuration information")
	c.Flags().BoolVarP(&rfconfig, "open", "o", false, "Open TOML configuration file")
	return c
}

func ConfigCmd(cmd *cobra.Command, args []string) {
	if show {
		c := config.DecodeTOMLConfig()
		fmt.Println("==========grpm Configuration==========")
		fmt.Println("Location:", "\033[1m", c.Location, "\033[0m", "=> Main grpm directory")
		fmt.Println("Downloaded:", "\033[1m", c.Downloaded, "\033[0m", "=> Where your downloaded files are saved")
		fmt.Println("library location:", "\033[1m", c.Library, "\033[0m", "=> Where your installed files are saved (files downloaded with --setup)")
		fmt.Println("System Architecture:", "\033[1m", c.Arch, "\033[0m", "=> Your system architecture (ex: x64,amd64,aarch64,etc)")
		fmt.Println("Operating System:", "\033[1m", c.Os, "\033[0m", "=> Your own operating system")

	} else if rfconfig {
		config.OpenTOMLConfig()
	} else {
		if err := cmd.Help(); err != nil {
			log.Fatal("Can't print config help command,", err)
		}
	}
}
