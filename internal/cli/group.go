package cli

import (
	"portal/internal/services"

	"github.com/spf13/cobra"
)

var groupCmd = &cobra.Command{
	Use:   "group",
	Short: "Manage repository groups",
	Long:  `Create, update, remove, and list groups of repositories.`,
}

var groupAddCmd = &cobra.Command{
	Use:   "add <group-name> <repo1> [repo2] ...",
	Short: "Create or update a group",
	Long:  `Create a new group or update an existing one with the specified repositories.`,
	Args:  cobra.MinimumNArgs(2),
	Run:   services.GroupAdd,
}

var groupRemoveCmd = &cobra.Command{
	Use:   "remove <group-name>",
	Short: "Remove a group",
	Long:  `Delete a group. This does not affect the repositories themselves.`,
	Args:  cobra.ExactArgs(1),
	Run:   services.GroupRemove,
}

var groupListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all groups",
	Long:  `Display all groups and their repositories.`,
	Run:   services.GroupList,
}

func init() {
	groupCmd.AddCommand(groupAddCmd)
	groupCmd.AddCommand(groupRemoveCmd)
	groupCmd.AddCommand(groupListCmd)
}
