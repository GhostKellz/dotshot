package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show what would be synced and git status",
	Run: func(cmd *cobra.Command, args []string) {
		showSyncPreview()
		showGitStatus()
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}

func showSyncPreview() {
	repoRoot := viper.GetString("repo_root")
	dotfiles := viper.GetStringMap("dotfiles")
	fmt.Println("Dotfiles sync preview:")
	for name, v := range dotfiles {
		m, ok := v.(map[string]interface{})
		if !ok {
			continue
		}
		paths, ok := m["paths"].([]interface{})
		if !ok {
			continue
		}
		for _, p := range paths {
			path := os.ExpandEnv(p.(string))
			dest := filepath.Join(repoRoot, name, filepath.Base(path))
			fmt.Printf("  %s -> %s\n", path, dest)
		}
	}
}

func showGitStatus() {
	repoRoot := viper.GetString("repo_root")
	cmd := exec.Command("git", "status")
	cmd.Dir = repoRoot
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Git status error: %v\n", err)
	}
	fmt.Println(string(out))
}
