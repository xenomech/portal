package services

import (
	"fmt"
	"portal/internal/config"
	"portal/internal/executor"
	"portal/internal/git"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func Pull(cmd *cobra.Command, args []string) {
	pullGroup, _ := cmd.Flags().GetString("group")
	pullRepo, _ := cmd.Flags().GetString("repo")

	cfg, err := config.Load()
	if err != nil {
		color.Red("Error loading config: %v", err)
		return
	}

	// Validate flag combinations
	if pullGroup != "" && pullRepo != "" {
		color.Red("Error: -g and -r flags are mutually exclusive")
		return
	}

	// Determine which repos to operate on
	repos, err := getTargetRepos(cfg, pullRepo, pullGroup, nil)
	if err != nil {
		color.Red("Error: %v", err)
		return
	}

	if len(repos) == 0 {
		fmt.Println("No repositories to operate on.")
		return
	}

	fmt.Printf("Pulling changes in %d repos...\n\n", len(repos))

	results := executor.Execute(repos, func(repo config.GitRepository) executor.Result {
		// Get current branch
		branch, err := git.GetCurrentBranch(repo.Path)
		if err != nil {
			return executor.Result{
				Repo:    repo,
				Success: false,
				Error:   fmt.Errorf("failed to get current branch: %w", err),
			}
		}

		// Pull from remote
		if err := git.Pull(repo.Path); err != nil {
			return executor.Result{
				Repo:    repo,
				Success: false,
				Error:   err,
			}
		}

		return executor.Result{
			Repo:    repo,
			Success: true,
			Message: fmt.Sprintf("→ %s (pulled)", branch),
		}
	})

	executor.PrintResults(results)

	successCount := executor.CountSuccess(results)
	fmt.Printf("\n%d/%d successful\n", successCount, len(results))
}
