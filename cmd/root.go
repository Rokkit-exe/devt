/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "devt",
	Short: "A CLI tool for developers to generate/encode/decode/hash",
	Long: `A CLI tool for developers to generate/encode/decode/hash.
	This tool provides various functionalities such as:
	- Generate random strings or objects
	- Encode and decode strings using different encoding types
	- Hash strings or files using different hash algorithms`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
