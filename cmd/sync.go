package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var dryRun bool

var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Sync dotfiles to repo",
	Run: func(cmd *cobra.Command, args []string) {
		SyncDotfiles()
		if viper.GetBool("git.auto_commit") {
			// GitAutoCommit function removed
		}
	},
}

func init() {
	rootCmd.AddCommand(syncCmd)
	syncCmd.Flags().BoolVar(&dryRun, "dry-run", false, "Show what would sync without copying files")
}
