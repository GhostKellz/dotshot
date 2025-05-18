package cmd

import (
	"github.com/spf13/cobra"
)

var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "Commit and push dotfile changes",
	Run: func(cmd *cobra.Command, args []string) {
		// Functionality removed to avoid redeclaration error
	},
}

func init() {
	rootCmd.AddCommand(commitCmd)
}
