package services

import (
	"fmt"
	"portal/internal/config"
	"portal/internal/executor"
	"portal/internal/git"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func Push(cmd *cobra.Command, args []string) {
	pushGroup, _ := cmd.Flags().GetString("group")
	pushRepo, _ := cmd.Flags().GetString("repo")
	setUpstream, _ := cmd.Flags().GetBool("set-upstream")

	cfg, err := config.Load()
	if err != nil {
		color.Red("Error loading config: %v", err)
		return
	}

	// Validate flag combinations
	if pushGroup != "" && pushRepo != "" {
		color.Red("Error: -g and -r flags are mutually exclusive")
		return
	}

	// Determine which repos to operate on
	repos, err := getTargetRepos(cfg, pushRepo, pushGroup, nil)
	if err != nil {
		color.Red("Error: %v", err)
		return
	}

	if len(repos) == 0 {
		fmt.Println("No repositories to operate on.")
		return
	}

	fmt.Printf("Pushing changes in %d repos...\n\n", len(repos))

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

		// Push to remote
		if err := git.Push(repo.Path, setUpstream); err != nil {
			return executor.Result{
				Repo:    repo,
				Success: false,
				Error:   err,
			}
		}

		message := fmt.Sprintf("→ %s (pushed)", branch)
		if setUpstream {
			message = fmt.Sprintf("→ %s (pushed with upstream)", branch)
		}

		return executor.Result{
			Repo:    repo,
			Success: true,
			Message: message,
		}
	})

	executor.PrintResults(results)

	successCount := executor.CountSuccess(results)
	fmt.Printf("\n%d/%d successful\n", successCount, len(results))
}
