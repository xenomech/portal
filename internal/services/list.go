package services

import (
	"fmt"
	"portal/internal/config"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func List(cmd *cobra.Command, args []string) {
	cfg, err := config.Load()
	if err != nil {
		color.Red("Error loading config: %v", err)
		return
	}

	repos := cfg.GetAllGitRepositories()
	if len(repos) == 0 {
		fmt.Println("No repositories registered. Use 'portal add <path>' to add one.")
		return
	}

	cyan := color.New(color.FgCyan).SprintFunc()
	dim := color.New(color.Faint).SprintFunc()

	// Find max name length for alignment
	maxLen := 0
	for _, r := range repos {
		if len(r.Name) > maxLen {
			maxLen = len(r.Name)
		}
	}

	fmt.Println()
	for _, r := range repos {
		name := fmt.Sprintf("%-*s", maxLen, r.Name)
		fmt.Printf("  %s  %s\n", cyan(name), dim(r.Path))
	}
	fmt.Println()
}
