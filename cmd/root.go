package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:  "quickshare",
	Long: `Quickshare is a cli tool for quickly sharing the files across the network`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(shareCmd)
	rootCmd.AddCommand(pairCmd)
	rootCmd.AddCommand(startCmd)
	rootCmd.AddCommand(acceptCmd)
}
