package cli

import (
	"portal/internal/services"

	"github.com/spf13/cobra"
)

var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Push changes to remote across repos, a group, or a single repo",
	Long:  `Push changes to the remote repository across all registered repositories, a specific group, or a single repo.`,
	Run:   services.Push,
}

func init() {
	pushCmd.Flags().StringP("group", "g", "", "Only operate on repos in this group")
	pushCmd.Flags().StringP("repo", "r", "", "Only operate on this specific repo")
	pushCmd.Flags().BoolP("set-upstream", "u", false, "Set upstream tracking for the current branch")
}
