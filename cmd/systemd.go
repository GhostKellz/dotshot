package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var systemdCmd = &cobra.Command{
	Use:   "systemd-helper",
	Short: "Print instructions for enabling dotshot as a systemd user service",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("To enable dotshot as a user service:")
		fmt.Println("1. Copy dotshot.service to ~/.config/systemd/user/dotshot.service")
		fmt.Println("2. Run: systemctl --user daemon-reload")
		fmt.Println("3. Run: systemctl --user enable --now dotshot.service")
		fmt.Println("4. Check status: systemctl --user status dotshot.service")
	},
}

func init() {
	rootCmd.AddCommand(systemdCmd)
}
