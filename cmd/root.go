package cmd

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "dotshot",
	Short: "Snapshot and sync your dotfiles with ease",
	Long:  `dotshot is a CLI tool for managing and syncing dotfiles across systems.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to dotshot ✨")
		fmt.Println("Use `dotshot sync`, `dotshot add`, or `dotshot init` to begin managing your dotfiles.")
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	if runtime.GOOS != "linux" {
		fmt.Println("Warning: dotshot is designed for Linux/Arch. Some features may not work on other OSes.")
	}

	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ./config.yaml)")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(".")
	}

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
