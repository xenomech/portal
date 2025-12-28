package services

import (
	"fmt"
	"portal/internal/config"
	"portal/internal/executor"
	"portal/internal/git"
	"portal/internal/utils"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func Checkout(cmd *cobra.Command, args []string) {
	// Get flags
	createBranch, _ := cmd.Flags().GetBool("create")
	checkoutGroup, _ := cmd.Flags().GetString("group")
	checkoutRepo, _ := cmd.Flags().GetString("repo")
	checkoutFromSpec, _ := cmd.Flags().GetStringArray("from")
	checkoutSync, _ := cmd.Flags().GetBool("sync")

	// Get branch name from args or --from flag
	var checkoutBranch string
	if len(args) > 0 {
		checkoutBranch = args[0]
	}

	if checkoutBranch == "" && len(checkoutFromSpec) == 0 {
		color.Red("Error: branch name is required")
		return
	}

	cfg, err := config.Load()
	if err != nil {
		color.Red("Error loading config: %v", err)
		return
	}

	// Parse --from specifications
	fromSpecs, err := utils.ParseFromSpecs(checkoutFromSpec)
	if err != nil {
		color.Red("Error: %v", err)
		return
	}

	// Validate flag combinations
	if len(fromSpecs) > 0 && (checkoutGroup != "" || checkoutRepo != "") {
		color.Red("Error: --from cannot be combined with -g or -r flags")
		return
	}
	if checkoutGroup != "" && checkoutRepo != "" {
		color.Red("Error: -g and -r flags are mutually exclusive")
		return
	}

	// Determine which repos to operate on
	repos, err := getTargetRepos(cfg, checkoutRepo, checkoutGroup, fromSpecs)
	if err != nil {
		color.Red("Error: %v", err)
		return
	}

	if len(repos) == 0 {
		fmt.Println("No repositories to operate on.")
		return
	}

	// Build descriptive message
	if len(fromSpecs) > 0 {
		fmt.Printf("Creating branch '%s' from specified base branches in %d repos...\n\n", checkoutBranch, len(repos))
	} else if createBranch {
		fmt.Printf("Creating and checking out branch '%s' in %d repos...\n\n", checkoutBranch, len(repos))
	} else {
		fmt.Printf("Checking out branch '%s' in %d repos...\n\n", checkoutBranch, len(repos))
	}

	results := executor.Execute(repos, func(repo config.GitRepository) executor.Result {
		// Step 1: Optionally sync (fetch) if --sync flag is set
		if checkoutSync {
			if err := git.Fetch(repo.Path); err != nil {
				return executor.Result{
					Repo:    repo,
					Success: false,
					Error:   fmt.Errorf("fetch failed: %w", err),
				}
			}
		}

		// Step 2: Check if branch already exists
		branchExists := git.BranchExists(repo.Path, checkoutBranch)

		// If not creating and branch doesn't exist, error
		if !createBranch && !branchExists {
			return executor.Result{
				Repo:    repo,
				Success: false,
				Error:   fmt.Errorf("branch '%s' does not exist (use -b to create)", checkoutBranch),
			}
		}

		// If branch exists and we're not forcing creation, just checkout
		if branchExists && !createBranch {
			if err := git.Checkout(repo.Path, checkoutBranch, false, ""); err != nil {
				return executor.Result{
					Repo:    repo,
					Success: false,
					Error:   err,
				}
			}
			return executor.Result{
				Repo:    repo,
				Success: true,
				Message: fmt.Sprintf("→ %s", checkoutBranch),
			}
		}

		// If branch exists and we're trying to create, just checkout existing
		if branchExists && createBranch {
			if err := git.Checkout(repo.Path, checkoutBranch, false, ""); err != nil {
				return executor.Result{
					Repo:    repo,
					Success: false,
					Error:   err,
				}
			}
			return executor.Result{
				Repo:    repo,
				Success: true,
				Message: fmt.Sprintf("→ %s (already exists)", checkoutBranch),
			}
		}

		// Step 3: Create branch - with or without base branch
		baseBranch, hasBase := fromSpecs[repo.Name]

		if hasBase {
			// Create from specified base branch
			if err := git.Checkout(repo.Path, checkoutBranch, true, baseBranch); err != nil {
				return executor.Result{
					Repo:    repo,
					Success: false,
					Error:   err,
				}
			}
			return executor.Result{
				Repo:    repo,
				Success: true,
				Message: fmt.Sprintf("→ %s (from %s)", checkoutBranch, baseBranch),
			}
		}

		// Create from current HEAD (existing behavior)
		if err := git.Checkout(repo.Path, checkoutBranch, true, ""); err != nil {
			return executor.Result{
				Repo:    repo,
				Success: false,
				Error:   err,
			}
		}
		return executor.Result{
			Repo:    repo,
			Success: true,
			Message: fmt.Sprintf("→ %s (created)", checkoutBranch),
		}
	})

	executor.PrintResults(results)

	// Show tree display for --from branches
	if len(fromSpecs) > 0 {
		utils.PrintBranchTree(results, fromSpecs, checkoutBranch)
	}

	successCount := executor.CountSuccess(results)
	fmt.Printf("\n%d/%d successful\n", successCount, len(results))
}

func getTargetRepos(cfg *config.Config, checkoutRepo, checkoutGroup string, fromSpecs map[string]string) ([]config.GitRepository, error) {
	var repos []config.GitRepository

	if len(fromSpecs) > 0 {
		// Mode: Only operate on repos specified in --from flags
		for repoName := range fromSpecs {
			repo, err := cfg.GetGitRepository(repoName)
			if err != nil {
				return nil, fmt.Errorf("repository '%s' not found in config", repoName)
			}
			repos = append(repos, *repo)
		}
	} else if checkoutRepo != "" {
		repo, err := cfg.GetGitRepository(checkoutRepo)
		if err != nil {
			return nil, err
		}
		repos = []config.GitRepository{*repo}
	} else if checkoutGroup != "" {
		var err error
		repos, err = cfg.GetGitRepositoryByGroup(checkoutGroup)
		if err != nil {
			return nil, err
		}
	} else {
		repos = cfg.GetAllGitRepositories()
	}

	return repos, nil
}
