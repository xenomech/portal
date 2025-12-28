package utils

import (
	"fmt"
	"portal/internal/executor"

	"github.com/fatih/color"
)

// PrintBranchTree displays a tree showing which branch the new branch was cut from
func PrintBranchTree(results []executor.Result, fromSpecs map[string]string, newBranch string) {
	green := color.New(color.FgGreen).SprintFunc()
	cyan := color.New(color.FgCyan).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()

	fmt.Println("\nBranch tree:")

	successfulResults := make([]executor.Result, 0)
	for _, r := range results {
		if r.Success {
			successfulResults = append(successfulResults, r)
		}
	}

	for i, r := range successfulResults {
		baseBranch := fromSpecs[r.Repo.Name]
		isLast := i == len(successfulResults)-1

		// Repo line
		if isLast {
			fmt.Printf("└── %s\n", yellow(r.Repo.Name))
		} else {
			fmt.Printf("├── %s\n", yellow(r.Repo.Name))
		}

		// Branch tree under repo
		prefix := "│   "
		if isLast {
			prefix = "    "
		}

		fmt.Printf("%s└── %s\n", prefix, cyan(baseBranch))
		fmt.Printf("%s    └── %s\n", prefix, green(newBranch))
	}
}
