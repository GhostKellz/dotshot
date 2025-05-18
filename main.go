package main

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func main() {
	// Load configuration
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("No config file found. You can run `dotshot init` to generate one.")
	}

	// Placeholder for CLI logic
	fmt.Println("Welcome to dotshot ✨")
	fmt.Println("Use `dotshot sync`, `dotshot add`, or `dotshot init` to begin managing your dotfiles.")

	os.Exit(0)
}

