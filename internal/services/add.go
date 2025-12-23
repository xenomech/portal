package services

import (
	"fmt"
	"os"
	"path/filepath"
	"portal/internal/config"
	"portal/internal/git"
	"portal/internal/utils"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func Add(cmd *cobra.Command, args []string) {
	path := args[0]

	absPath, err := utils.GetAbsolutePath(path)
	if err != nil {
		color.Red("Error resolving absolute path: %s", err)
		os.Exit(1)
	}

	if !utils.DoesPathExist(absPath) {
		color.Red("Path does not exist: %s", absPath)
		os.Exit(1)
	}

	if !git.IsGitRepo(absPath) {
		color.Red("Not a git repository: %s", absPath)
		os.Exit(1)
	}

	name, err := cmd.Flags().GetString("name")
	if err != nil {
		color.Red("Error getting name flag: %v", err)
		os.Exit(1)
	}
	if name == "" {
		name = filepath.Base(absPath)
	}

	cfg, err := config.Load()
	if err != nil {
		color.Red("Error loading config: %v", err)
		os.Exit(1)
	}
	if err := cfg.AddGitRepositoryToConfig(absPath, name); err != nil {
		color.Red("Error: %v", err)
		os.Exit(1)
	}
	if err := cfg.Save(); err != nil {
		color.Red("Error saving config: %v", err)
		os.Exit(1)
	}

	green := color.New(color.FgGreen).SprintFunc()
	cyan := color.New(color.FgCyan).SprintFunc()
	fmt.Printf("%s Added %s (%s)\n", green("✓"), cyan(name), absPath)

}
