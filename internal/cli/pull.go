package cli

import (
	"portal/internal/services"

	"github.com/spf13/cobra"
)

var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Pull changes from remote across repos, a group, or a single repo",
	Long:  `Pull changes from the remote repository across all registered repositories, a specific group, or a single repo.`,
	Run:   services.Pull,
}

func init() {
	pullCmd.Flags().StringP("group", "g", "", "Only operate on repos in this group")
	pullCmd.Flags().StringP("repo", "r", "", "Only operate on this specific repo")
}
