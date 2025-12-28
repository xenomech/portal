package cli

import (
	"portal/internal/services"

	"github.com/spf13/cobra"
)

var checkoutCmd = &cobra.Command{
	Use:   "checkout <branch-name>",
	Short: "Checkout a branch across repos, a group, or a single repo",
	Long:  `Checkout an existing branch or create a new one with -b flag across all registered repositories, a specific group, or a single repo.`,
	Args:  cobra.ExactArgs(1),
	Run:   services.Checkout,
}

func init() {
	checkoutCmd.Flags().BoolP("create", "b", false, "Create a new branch")
	checkoutCmd.Flags().StringP("group", "g", "", "Only operate on repos in this group")
	checkoutCmd.Flags().StringP("repo", "r", "", "Only operate on this specific repo")
	checkoutCmd.Flags().StringArrayP("from", "f", nil, "Base branch for specific repo (format: repo=branch, can be repeated)")
	checkoutCmd.Flags().Bool("sync", false, "Fetch from remote before creating branch")
}
