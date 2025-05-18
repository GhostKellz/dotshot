package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/viper"
)

// GitAutoCommit commits and pushes changes if enabled
func GitAutoCommit() {
	repoRoot := viper.GetString("repo_root")
	msg := viper.GetString("git.commit_message")
	cmds := [][]string{
		{"git", "add", "."},
		{"git", "commit", "-m", msg},
		{"git", "push"},
	}
	for _, args := range cmds {
		cmd := exec.Command(args[0], args[1:]...)
		cmd.Dir = repoRoot
		out, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Printf("Git error: %v\n%s\n", err, string(out))
		}
	}
	fmt.Println("Git auto-commit and push complete.")
}
