package main_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/ghostkellz/dotshot/cmd"
	"github.com/spf13/viper"
)

func TestSyncDotfiles(t *testing.T) {
	os.Setenv("HOME", "/tmp")
	viper.Set("repo_root", "/tmp/dotshot-test-repo")
	viper.Set("dotfiles", map[string]interface{}{
		"test": map[string]interface{}{
			"paths": []interface{}{filepath.Join(os.TempDir(), "testfile")},
		},
	})
	f, err := os.Create(filepath.Join(os.TempDir(), "testfile"))
	if err != nil {
		t.Fatal(err)
	}
	f.WriteString("hello")
	f.Close()
	cmd.SyncDotfiles()
	if _, err := os.Stat(filepath.Join("/tmp/dotshot-test-repo", "test", "testfile")); err != nil {
		t.Errorf("File not synced: %v", err)
	}
}
