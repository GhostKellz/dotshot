package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var restoreCmd = &cobra.Command{
	Use:   "restore",
	Short: "Restore dotfiles from repo to original locations",
	Run: func(cmd *cobra.Command, args []string) {
		repoRoot := viper.GetString("repo_root")
		dotfiles := viper.GetStringMap("dotfiles")
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
				orig := os.ExpandEnv(p.(string))
				src := filepath.Join(repoRoot, name, filepath.Base(orig))
				copyPath(src, orig)
			}
		}
		fmt.Println("Dotfiles restored from repo.")
	},
}

func init() {
	rootCmd.AddCommand(restoreCmd)
}
