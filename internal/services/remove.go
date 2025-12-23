package services

import (
	"fmt"
	"os"
	"portal/internal/config"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func Remove(cmd *cobra.Command, args []string) {
	name := args[0]

	cfg, err := config.Load()
	if err != nil {
		color.Red("Error loading config: %v", err)
		os.Exit(1)
	}

	if err := cfg.RemoveGitRepositoryFromConfig(name); err != nil {
		color.Red("Error: %v", err)
		os.Exit(1)
	}

	if err := cfg.Save(); err != nil {
		color.Red("Error saving config: %v", err)
		os.Exit(1)
	}

	green := color.New(color.FgGreen).SprintFunc()
	cyan := color.New(color.FgCyan).SprintFunc()
	fmt.Printf("%s Removed %s\n", green("✓"), cyan(name))
}
