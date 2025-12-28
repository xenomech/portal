package services

import (
	"fmt"
	"portal/internal/config"
	"portal/internal/executor"
	"portal/internal/git"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func Switch(cmd *cobra.Command, args []string) {
	branchName := args[0]
	switchGroup, _ := cmd.Flags().GetString("group")
	switchRepo, _ := cmd.Flags().GetString("repo")

	cfg, err := config.Load()
	if err != nil {
		color.Red("Error loading config: %v", err)
		return
	}

	// Validate flag combinations
	if switchGroup != "" && switchRepo != "" {
		color.Red("Error: -g and -r flags are mutually exclusive")
		return
	}

	// Determine which repos to operate on
	repos, err := getTargetRepos(cfg, switchRepo, switchGroup, nil)
	if err != nil {
		color.Red("Error: %v", err)
		return
	}

	if len(repos) == 0 {
		fmt.Println("No repositories to operate on.")
		return
	}

	fmt.Printf("Switching to branch '%s' in %d repos...\n\n", branchName, len(repos))

	results := executor.Execute(repos, func(repo config.GitRepository) executor.Result {
		// Check if branch exists
		if !git.BranchExists(repo.Path, branchName) {
			return executor.Result{
				Repo:    repo,
				Success: false,
				Error:   fmt.Errorf("branch '%s' does not exist", branchName),
			}
		}

		// Switch to the branch
		if err := git.Checkout(repo.Path, branchName, false, ""); err != nil {
			return executor.Result{
				Repo:    repo,
				Success: false,
				Error:   err,
			}
		}

		return executor.Result{
			Repo:    repo,
			Success: true,
			Message: fmt.Sprintf("→ %s", branchName),
		}
	})

	executor.PrintResults(results)

	successCount := executor.CountSuccess(results)
	fmt.Printf("\n%d/%d successful\n", successCount, len(results))
}
