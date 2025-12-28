package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "portal",
	Short: "Manage git branches across multiple repositories",
	Long: fmt.Sprintf(`Portal (%s-%s)

Portal is a CLI tool for managing git branches across multiple repositories.
It allows you to checkout, switch, and view status across all your repos with a single command.`,
		VersionName, Version),
	Version: Version,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(removeCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(groupCmd)
	rootCmd.AddCommand(checkoutCmd)
	rootCmd.AddCommand(switchCmd)
	rootCmd.AddCommand(pullCmd)
	rootCmd.AddCommand(pushCmd)
}
