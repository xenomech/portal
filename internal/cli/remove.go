package cli

import (
	"portal/internal/services"

	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove <name>",
	Short: "Remove a repository from portal",
	Long:  `Unregister a repository from portal. This does not delete the actual repository.`,
	Args:  cobra.ExactArgs(1),
	Run:   services.Remove,
}
