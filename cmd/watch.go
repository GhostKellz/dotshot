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

var verbose bool

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
				addWatchRecursive(watcher, path)
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
	watchCmd.Flags().BoolVar(&verbose, "verbose", false, "Enable verbose output")
}

// getExcludes returns a map of excluded patterns from config
func getExcludes() map[string]struct{} {
	ex := map[string]struct{}{}
	exList := viper.GetStringSlice("exclude")
	for _, e := range exList {
		ex[e] = struct{}{}
	}
	return ex
}

func isExcluded(path string, excludes map[string]struct{}) bool {
	for ex := range excludes {
		if matched, _ := filepath.Match(ex, filepath.Base(path)); matched {
			return true
		}
	}
	return false
}

// SyncDotfiles is called by the watcher and sync command
func SyncDotfiles() {
	repoRoot := viper.GetString("repo_root")
	dotfiles := viper.GetStringMap("dotfiles")
	excludes := getExcludes()
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
			if isExcluded(path, excludes) {
				continue
			}
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
		if verbose {
			fmt.Printf("Error stating %s: %v\n", src, err)
		}
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
		if verbose {
			fmt.Printf("Error opening %s: %v\n", src, err)
		}
		return
	}
	defer in.Close()
	os.MkdirAll(filepath.Dir(dest), 0755)
	out, err := os.Create(dest)
	if err != nil {
		if verbose {
			fmt.Printf("Error creating %s: %v\n", dest, err)
		}
		return
	}
	defer out.Close()
	_, err = io.Copy(out, in)
	if err != nil && verbose {
		fmt.Printf("Error copying %s to %s: %v\n", src, dest, err)
	}
}

func copyDir(src, dest string) {
	entries, err := os.ReadDir(src)
	if err != nil {
		if verbose {
			fmt.Printf("Error reading dir %s: %v\n", src, err)
		}
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

// addWatchRecursive adds a path to the watcher recursively
func addWatchRecursive(watcher *fsnotify.Watcher, path string) error {
	info, err := os.Stat(path)
	if err != nil {
		return err
	}
	if !info.IsDir() {
		return watcher.Add(path)
	}
	// Add the directory itself
	if err := watcher.Add(path); err != nil {
		return err
	}
	entries, err := os.ReadDir(path)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		child := filepath.Join(path, entry.Name())
		if entry.IsDir() {
			addWatchRecursive(watcher, child)
		} else {
			watcher.Add(child)
		}
	}
	return nil
}
