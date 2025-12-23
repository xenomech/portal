package git

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

// TODO: need better logging
func runGit(repoPath string, args ...string) (string, error) {
	cmd := exec.Command("git", args...)
	cmd.Dir = repoPath
	fmt.Println(cmd.String())
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("%s: %s", err, stderr.String())
	}

	return strings.TrimSpace(stdout.String()), nil
}

// idiom, i only care about the error!!
// pass in a path, the code flows like... sets the PWD as path and runs `git rev-pars --git-dir` which should output .git if its a git repo
func IsGitRepo(path string) bool {
	_, err := runGit(path, "rev-parse", "--git-dir")
	return err == nil
}

/*
*
what it does is, tells you whcic is your current branch, reverse parse HEAD  will give a  commit hash, we need to know in which branch our head is so add a --abbrev-ref it will give teh branch
*
*/
func GetCurrentBranch(repoPath string) (string, error) {
	return runGit(repoPath, "rev-parse", "--abbrev-ref", "HEAD")
}

/*
*
what it does, is that it runs git checkout to a branch, if create is there  add -b to the args list, if a base branch is provided, we need to validate it and checkout from there
*
*/
func Checkout(repoPath, branch string, create bool, baseBranch string) error {
	args := []string{"checkout"}
	if create {
		args = append(args, "-b")
	}
	args = append(args, branch)

	if baseBranch != "" {
		if !BranchExists(repoPath, baseBranch) {
			return fmt.Errorf("base branch '%s' not found", baseBranch)
		}
		args = append(args, baseBranch)
	}

	_, err := runGit(repoPath, args...)
	return err
}

/*
*
what it does is, tells you if the branch exist or not, reverse parse branch  will give a  commit hash that the branch points to,and --verify is to make sure it resolves to a valid object
*
*/
func BranchExists(repoPath, branch string) bool {
	_, err := runGit(repoPath, "rev-parse", "--verify", branch)
	return err == nil
}

// Fetch fetches from the remote
func Fetch(repoPath string) error {
	_, err := runGit(repoPath, "fetch", "--quiet")
	return err
}

// Pull pulls from the remote
func Pull(repoPath string) error {
	_, err := runGit(repoPath, "pull", "--quiet")
	return err
}
