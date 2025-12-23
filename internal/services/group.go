package services

import (
	"fmt"
	"os"
	"portal/internal/config"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func GroupAdd(cmd *cobra.Command, args []string) {
	groupName := args[0]
	repoNames := args[1:]

	cfg, err := config.Load()
	if err != nil {
		color.Red("Error loading config: %v", err)
		os.Exit(1)
	}

	if err := cfg.AddGroup(groupName, repoNames); err != nil {
		color.Red("Error: %v", err)
		os.Exit(1)
	}

	if err := cfg.Save(); err != nil {
		color.Red("Error saving config: %v", err)
		os.Exit(1)
	}

	green := color.New(color.FgGreen).SprintFunc()
	cyan := color.New(color.FgCyan).SprintFunc()
	fmt.Printf("%s Group %s updated with %d repos\n", green("✓"), cyan(groupName), len(repoNames))
}

func GroupRemove(cmd *cobra.Command, args []string) {
	groupName := args[0]

	cfg, err := config.Load()
	if err != nil {
		color.Red("Error loading config: %v", err)
		os.Exit(1)
	}

	if err := cfg.RemoveGroup(groupName); err != nil {
		color.Red("Error: %v", err)
		os.Exit(1)
	}

	if err := cfg.Save(); err != nil {
		color.Red("Error saving config: %v", err)
		os.Exit(1)
	}

	green := color.New(color.FgGreen).SprintFunc()
	cyan := color.New(color.FgCyan).SprintFunc()
	fmt.Printf("%s Group %s removed\n", green("✓"), cyan(groupName))
}

func GroupList(cmd *cobra.Command, args []string) {
	cfg, err := config.Load()
	if err != nil {
		color.Red("Error loading config: %v", err)
		return
	}

	if len(cfg.Groups) == 0 {
		fmt.Println("No groups defined. Use 'portal group add <name> <repos...>' to create one.")
		return
	}

	cyan := color.New(color.FgCyan).SprintFunc()
	dim := color.New(color.Faint).SprintFunc()

	fmt.Println()
	for name, repos := range cfg.Groups {
		fmt.Printf("  %s\n", cyan(name))
		for _, repo := range repos {
			fmt.Printf("    %s %s\n", dim("•"), repo)
		}
	}
	fmt.Println()
}
