package cli

import (
	"portal/internal/services"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add <path>",
	Short: "Add a repository to portal",
	Long:  `Register a git repository with portal. The repository name is auto-detected from the folder name unless --name is specified.`,
	Args:  cobra.ExactArgs(1),
	Run:   services.Add,
}

func init() {
	addCmd.Flags().StringP("name", "n", "", "Custom name for the repository")
}
