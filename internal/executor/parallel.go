package executor

import (
	"fmt"
	"portal/internal/config"
	"sync"

	"github.com/fatih/color"
)

// Result represents the result of an operation on a single repo
type Result struct {
	Repo    config.GitRepository
	Success bool
	Message string
	Error   error
}

// Operation is a function that performs an operation on a repo
type Operation func(repo config.GitRepository) Result

// Execute runs an operation on all repos in parallel and returns results
func Execute(repos []config.GitRepository, op Operation) []Result {
	//  how this works, we need to create a pre allocated slice for results, and as the results slice is made with the length of the repos, each repo result will be in the index
	results := make([]Result, len(repos))
	// what sync waitgroup does is its used to manage lifecycel of goroutines
	var wg sync.WaitGroup

	// Loop through each repository with index and value
	for i, repo := range repos {
		// Increment WaitGroup counter by 1 before launching goroutine (tells WaitGroup "one more goroutine to wait for")
		wg.Add(1)
		// Launch a self invocating function as a goroutine
		go func(idx int, r config.GitRepository) {
			// WaitGroup decrement when function exits
			defer wg.Done()
			// Execute the operation function on this repository and store result at index :: thread-safe, each goroutine writes to unique index
			results[idx] = op(r)
		}(i, repo)
	}

	// Block execution until WaitGroup counter reaches zero ie, all goroutines have called Done()
	wg.Wait()
	return results
}

// PrintResults prints the results in a formatted way
func PrintResults(results []Result) {
	green := color.New(color.FgGreen).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()
	cyan := color.New(color.FgCyan).SprintFunc()

	// Find max repo name length for alignment
	maxLen := 0
	for _, r := range results {
		if len(r.Repo.Name) > maxLen {
			maxLen = len(r.Repo.Name)
		}
	}

	for _, r := range results {
		name := fmt.Sprintf("%-*s", maxLen, r.Repo.Name)
		if r.Success {
			fmt.Printf(" %s %s  %s\n", green("✓"), cyan(name), r.Message)
		} else {
			errMsg := r.Message
			if r.Error != nil {
				errMsg = r.Error.Error()
			}
			fmt.Printf(" %s %s  %s\n", red("✗"), cyan(name), red(errMsg))
		}
	}
}

func CountSuccess(results []Result) int {
	count := 0
	for _, r := range results {
		if r.Success {
			count++
		}
	}
	return count
}
