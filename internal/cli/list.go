package cli

import (
	"portal/internal/services"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all registered repositories",
	Long:  `Display all repositories registered with portal.`,
	Run:   services.List,
}
