package cli

import (
	"portal/internal/services"

	"github.com/spf13/cobra"
)

var switchCmd = &cobra.Command{
	Use:   "switch [branch-name]",
	Short: "Switch to an existing branch across repos, a group, or a single repo",
	Long:  `Switch to an existing branch across all registered repositories, a specific group, or a single repo.`,
	Args:  cobra.ExactArgs(1),
	Run:   services.Switch,
}

func init() {
	switchCmd.Flags().StringP("group", "g", "", "Only operate on repos in this group")
	switchCmd.Flags().StringP("repo", "r", "", "Only operate on this specific repo")
}
