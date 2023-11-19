package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func Execute() {
	rootCmd := &cobra.Command{}

	// Disable the completion command
	rootCmd.CompletionOptions.DisableDefaultCmd = true

	rootCmd.AddCommand(createService())
	rootCmd.AddCommand(createRoute())
	rootCmd.AddCommand(createView())
	rootCmd.AddCommand(createAllFiles())

	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("root command: %s", err.Error())
		os.Exit(1)
	}
}
