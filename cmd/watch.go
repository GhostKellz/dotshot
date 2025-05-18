package cmd

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var watchCmd = &cobra.Command{
	Use:   "watch",
	Short: "Watch dotfiles for changes and sync automatically",
	Run: func(cmd *cobra.Command, args []string) {
		watcher, err := fsnotify.NewWatcher()
		if err != nil {
			log.Fatal(err)
		}
		defer watcher.Close()

		dotfiles := viper.GetStringMap("dotfiles")
		for _, v := range dotfiles {
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
				if err := watcher.Add(path); err != nil {
					fmt.Printf("Failed to watch %s: %v\n", path, err)
				}
			}
		}

		fmt.Println("Watching dotfiles for changes...")
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				fmt.Printf("Change detected: %s\n", event.Name)
				SyncDotfiles()
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				fmt.Println("Watcher error:", err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(watchCmd)
	watchCmd.Flags().BoolVar(&dryRun, "dry-run", false, "Preview sync changes without applying them")
}

// SyncDotfiles is called by the watcher and sync command
func SyncDotfiles() {
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
			path := os.ExpandEnv(p.(string))
			dest := filepath.Join(repoRoot, name, filepath.Base(path))
			if dryRun {
				fmt.Printf("[dry-run] Would sync %s -> %s\n", path, dest)
			} else {
				copyPath(path, dest)
			}
		}
	}
	if dryRun {
		fmt.Println("[dry-run] Dotfiles sync preview complete.")
	} else {
		fmt.Println("Dotfiles synced.")
	}
}

// copyPath copies a file or directory
func copyPath(src, dest string) {
	info, err := os.Stat(src)
	if err != nil {
		fmt.Printf("Error stating %s: %v\n", src, err)
		return
	}
	if info.IsDir() {
		copyDir(src, dest)
	} else {
		copyFile(src, dest)
	}
}

func copyFile(src, dest string) {
	in, err := os.Open(src)
	if err != nil {
		fmt.Printf("Error opening %s: %v\n", src, err)
		return
	}
	defer in.Close()
	os.MkdirAll(filepath.Dir(dest), 0755)
	out, err := os.Create(dest)
	if err != nil {
		fmt.Printf("Error creating %s: %v\n", dest, err)
		return
	}
	defer out.Close()
	_, err = io.Copy(out, in)
	if err != nil {
		fmt.Printf("Error copying %s to %s: %v\n", src, dest, err)
	}
}

func copyDir(src, dest string) {
	entries, err := os.ReadDir(src)
	if err != nil {
		fmt.Printf("Error reading dir %s: %v\n", src, err)
		return
	}
	os.MkdirAll(dest, 0755)
	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		destPath := filepath.Join(dest, entry.Name())
		if entry.IsDir() {
			copyDir(srcPath, destPath)
		} else {
			copyFile(srcPath, destPath)
		}
	}
}
